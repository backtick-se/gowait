package daemon

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"
	"time"
)

type Worker interface {
	ID() task.ID
	Image() string
	Status() executor.Status
	Start(context.Context) error

	OnInit()
	OnStop()
	OnIdle()
	OnAquire(task.Instance)
}

// executor instance
type worker struct {
	id     task.ID
	driver cluster.Driver
	image  string
	status executor.Status

	on_init   chan struct{}
	on_idle   chan struct{}
	on_stop   chan struct{}
	on_aquire chan task.Instance
}

type TaskEventFn func(event string, state task.Run)

func NewWorker(driver cluster.Driver, id task.ID, image string) Worker {
	t := &worker{
		id:     id,
		driver: driver,
		image:  image,
		status: executor.StatusWait,

		on_init:   make(chan struct{}),
		on_idle:   make(chan struct{}),
		on_stop:   make(chan struct{}),
		on_aquire: make(chan task.Instance),
	}
	return t
}

func (w *worker) ID() task.ID   { return w.id }
func (w *worker) Image() string { return w.image }

func (w *worker) Status() executor.Status { return w.status }

func (w *worker) OnInit() { w.on_init <- struct{}{} }
func (w *worker) OnIdle() { w.on_idle <- struct{}{} }
func (w *worker) OnStop() { w.on_stop <- struct{}{} }

func (w *worker) OnAquire(i task.Instance) {
	w.on_aquire <- i
}

func (i *worker) Start(ctx context.Context) error {
	if err := i.driver.Spawn(ctx, i.id, i.image); err != nil {
		return fmt.Errorf("failed to spawn task %s: %w", i.id, err)
	}
	go i.proc()
	return nil
}

func (i *worker) proc() {
	defer i.cleanup()

	for {
		select {
		case <-i.on_init:
			i.status = executor.StatusIdle
		case <-i.on_idle:
			i.status = executor.StatusIdle
		case <-i.on_stop:
			i.status = executor.StatusStop
			return
		case instance := <-i.on_aquire:
			i.process_next(instance)
		}
	}
}

func (i *worker) process_next(instance task.Instance) {
	done := make(chan struct{})
	fmt.Println("dequeued", instance)

	// change state to Executing
	i.status = executor.StatusExec

	// launch a goroutine to handle instance events
	go instance.Exec(done)

	// wait for execution to complete
	for {
		select {
		case <-done:
			i.status = executor.StatusIdle
			return
		case <-i.on_stop:
			i.status = executor.StatusStop
			return
		case <-time.After(10 * time.Second):
			// periodic liveness check
			fmt.Println("poke", i.id)
			ctx := context.Background()
			if err := i.driver.Poke(ctx, i.id); err != nil {
				// crash detected.
				// maybe we want to handle crashes differently? e.g. re-try task
				fmt.Println("executor", i.id, "failed liveness check:", err)
				instance.OnFailure(&task.MsgFailure{
					Header: task.Header{
						ID: string(instance.ID()),
					},
					Error: fmt.Errorf("cluster task error: %w", err),
				})
				i.status = executor.StatusCrash
				return
			}
		}
	}

	// change state to Idle
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
