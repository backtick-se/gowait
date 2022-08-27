package daemon

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"
	"time"
)

// executor instance
type worker struct {
	id     task.ID
	driver cluster.Driver
	image  string

	publish    TaskEventFn
	on_dequeue chan *instance
}

type TaskEventFn func(event string, state task.Run)

func newWorker(driver cluster.Driver, id task.ID, image string, callback TaskEventFn) *worker {
	t := &worker{
		id:     id,
		driver: driver,
		image:  image,

		publish:    callback,
		on_dequeue: make(chan *instance),
	}
	go t.proc()
	return t
}

func (i *worker) proc() {
	defer i.cleanup()

	// this is the instance management loop
	// at this point the task is in the "scheduled" state
	// i suppose we start by calling cluster.Spawn() ?
	if err := i.driver.Spawn(context.Background(), i.id, i.image); err != nil {
		fmt.Println("failed to spawn task", i.id, ":", err)
		return
	}

	// todo: this should be structured as a finite state machine

	for {
		i.aquire()
	}
}

func (i *worker) aquire() {
	instance := <-i.on_dequeue
	done := make(chan struct{})
	fmt.Println("dequeued", instance)

	go instance.exec(done)

	for {
		select {
		case <-done:
			return
		case <-time.After(10 * time.Second):
			// periodic liveness check
			fmt.Println("poke", i.id)
			ctx := context.Background()
			if err := i.driver.Poke(ctx, i.id); err != nil {
				fmt.Println("executor", i.id, "failed liveness check:", err)
				instance.State().Fail(fmt.Errorf("cluster task error: %w", err))
				return
			}
		}
	}
}

func (i *worker) cleanup() {

	// wait a sec for any logs to arrive
	// todo: avoid race condition here
	time.Sleep(time.Second)

	// delete executor pod
	// ctx := context.Background()
	// if err := i.driver.Kill(ctx, i.id); err != nil {
	// 	// log error
	// 	fmt.Println("failed to kill", i, ":", err)
	// }
}
