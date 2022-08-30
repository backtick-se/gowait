package daemon

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"
	"go.uber.org/zap"

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
	driver cluster.Driver
	log    *zap.Logger

	id     task.ID
	image  string
	status executor.Status
	task   task.Instance

	on_init   chan struct{}
	on_stop   chan struct{}
	on_aquire chan task.Instance
	err       chan error
}

func NewWorker(driver cluster.Driver, id task.ID, image string, log *zap.Logger) Worker {
	t := &worker{
		driver: driver,
		log:    log.With(zap.String("executor", string(id))),

		id:     id,
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
	w.log.Info("executor idle")

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
	w.log.Info("executor aquire", zap.String("task", string(w.task.ID())))

	for {
		select {
		case <-done:
			return w.state_idle, nil

		case <-w.on_stop:
			// this might be a reasonable place to fail the running task
			return w.state_stop, nil

		case <-time.After(10 * time.Second):
			// periodic liveness check
			w.log.Debug("executor poke", zap.String("task", string(w.task.ID())))
			ctx := context.Background()
			if err := w.driver.Poke(ctx, w.id); err != nil {
				// crash detected.
				// maybe we want to handle crashes differently? e.g. re-try task
				w.log.Error("executor failed liveness check", zap.String("task", string(w.task.ID())), zap.Error(err))
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
	w.log.Info("executor stop")

	return nil, nil
}

func (w *worker) cleanup() {
	w.log.Debug("executor cleanup")

	// wait a sec for any logs to arrive
	// todo: avoid race condition here
	time.Sleep(time.Second)

	// delete executor pod
	ctx := context.Background()
	if err := w.driver.Kill(ctx, w.id); err != nil {
		// log error
		w.log.Debug("executor cleanup failed", zap.Error(err))
	}
}
