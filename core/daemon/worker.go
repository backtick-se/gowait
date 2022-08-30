package daemon

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"errors"
	"fmt"
	"time"
)

var ErrInvalidState = errors.New("invalid state")

type stateFn func() (stateFn, error)

type Worker interface {
	ID() task.ID
	Image() string
	Status() executor.Status
	Start(context.Context) error

	OnInit() error
	OnStop() error
	OnAquire(task.Instance) error
}

// executor instance
type worker struct {
	id     task.ID
	driver cluster.Driver
	image  string
	status executor.Status

	task task.Instance

	on_init   chan struct{}
	on_stop   chan struct{}
	on_aquire chan task.Instance
	err       chan error
}

func NewWorker(driver cluster.Driver, id task.ID, image string) Worker {
	t := &worker{
		id:     id,
		driver: driver,
		image:  image,
		status: executor.StatusWait,

		on_init:   make(chan struct{}),
		on_stop:   make(chan struct{}),
		on_aquire: make(chan task.Instance),
		err:       make(chan error),
	}
	return t
}

func (w *worker) ID() task.ID   { return w.id }
func (w *worker) Image() string { return w.image }

func (w *worker) Status() executor.Status { return w.status }

func (w *worker) OnInit() error { w.on_init <- struct{}{}; return <-w.err }
func (w *worker) OnStop() error { w.on_stop <- struct{}{}; return <-w.err }

func (w *worker) OnAquire(i task.Instance) error { w.on_aquire <- i; return <-w.err }

func (w *worker) Start(ctx context.Context) error {
	if err := w.driver.Spawn(ctx, w.id, w.image); err != nil {
		return fmt.Errorf("failed to spawn task %s: %w", w.id, err)
	}
	go w.proc()
	return nil
}

func (w *worker) proc() {
	defer w.cleanup()

	var err error
	state := w.state_wait
	for {
		state, err = state()
		if err != nil {
			w.task = nil
			w.status = executor.StatusCrash
			return
		}
		if state == nil {
			return
		}
	}
}

// Acknowledge state transfer
func (w *worker) ack() {
	select {
	case w.err <- nil:
	case <-time.After(time.Millisecond):
	}
}

func (w *worker) state_wait() (stateFn, error) {
	for {
		select {
		case <-w.on_init:
			return w.state_idle, nil
		case <-w.on_stop:
			return w.state_stop, nil

		// invalid transitions:
		case <-w.on_aquire:
			w.err <- ErrInvalidState
		}
	}
}

func (w *worker) state_idle() (stateFn, error) {
	w.task = nil
	w.status = executor.StatusIdle
	fmt.Println("executor", w.id, "idle")

	w.ack()

	for {
		select {
		case w.task = <-w.on_aquire:
			return w.state_exec, nil
		case <-w.on_stop:
			return w.state_stop, nil

		// invalid transitions:
		case <-w.on_init:
			w.err <- ErrInvalidState
		}
	}
}

func (w *worker) state_exec() (stateFn, error) {
	if w.task == nil {
		return nil, fmt.Errorf("no worker task set")
	}
	w.status = executor.StatusExec

	// launch a goroutine to handle instance events
	// maybe its weird that this happens here.
	done := w.task.Start()

	w.ack()
	fmt.Println("executor", w.id, "running", w.task.ID())

	for {
		select {
		case <-done:
			return w.state_idle, nil

		case <-w.on_stop:
			// this might be a reasonable place to fail the running task
			return w.state_stop, nil

		case <-time.After(10 * time.Second):
			// periodic liveness check
			fmt.Println("poke", w.id)
			ctx := context.Background()
			if err := w.driver.Poke(ctx, w.id); err != nil {
				// crash detected.
				// maybe we want to handle crashes differently? e.g. re-try task
				fmt.Println("executor", w.id, "failed liveness check:", err)
				w.task.Fail(err)
				return nil, err
			}

		// invalid transitions:
		case <-w.on_aquire:
			w.err <- ErrInvalidState
		case <-w.on_init:
			w.err <- ErrInvalidState
		}
	}
}

func (w *worker) state_stop() (stateFn, error) {
	w.task = nil
	w.status = executor.StatusStop
	w.ack()
	fmt.Println("executor", w.id, "stopped")

	return nil, nil
}

func (i *worker) cleanup() {
	fmt.Println("cleanup executor", i.id)

	// wait a sec for any logs to arrive
	// todo: avoid race condition here
	time.Sleep(time.Second)

	// delete executor pod
	ctx := context.Background()
	if err := i.driver.Kill(ctx, i.id); err != nil {
		// log error
		fmt.Println("failed to kill", i, ":", err)
	}
}
