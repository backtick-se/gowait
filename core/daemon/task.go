package daemon

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/msg"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"
	"time"
)

type instance struct {
	id     task.ID
	driver cluster.Driver
	state  task.State
	logs   map[string][]string

	publish     TaskEventFn
	on_init     chan *msg.TaskInit
	on_fail     chan *msg.TaskFailure
	on_complete chan *msg.TaskComplete
	on_log      chan *msg.LogEntry
}

var _ task.T = &instance{}

type TaskEventFn func(event string, state task.State)

func newInstance(driver cluster.Driver, id task.ID, spec *task.Spec, callback TaskEventFn) *instance {
	t := &instance{
		id:     id,
		driver: driver,
		state: task.State{
			ID:        id,
			Spec:      spec,
			Status:    task.StatusWait,
			Scheduled: time.Now(),
		},
		logs: make(map[string][]string),

		publish:     callback,
		on_init:     make(chan *msg.TaskInit),
		on_fail:     make(chan *msg.TaskFailure),
		on_complete: make(chan *msg.TaskComplete),
		on_log:      make(chan *msg.LogEntry),
	}
	go t.proc()
	return t
}

func (i *instance) ID() task.ID       { return i.state.ID }
func (i *instance) Spec() *task.Spec  { return i.state.Spec }
func (i *instance) State() task.State { return i.state }

func (i *instance) Logs(file string) ([]string, error) {
	if logs, ok := i.logs[file]; ok {
		return logs, nil
	}
	return nil, fmt.Errorf("no such log: %s", file)
}

func (i *instance) proc() {
	defer i.cleanup()

	spec := i.state.Spec

	// a specific context for each task allows us to set per-task deadlines etc
	ctx := context.Background()

	if spec.Timeout > 0 {
		deadline, cancel := context.WithTimeout(ctx, time.Duration(spec.Timeout)*time.Second)
		defer cancel()
		ctx = deadline
	}

	i.publish("task/schedule", i.state)

	// this is the instance management loop
	// at this point the task is in the "scheduled" state
	// i suppose we start by calling cluster.Spawn() ?
	if err := i.driver.Spawn(ctx, i.id, spec); err != nil {
		fmt.Println("failed to spawn task", i.id, ":", err)
		return
	}

	// todo: this should be structured as a finite state machine

	for {
		select {
		case <-i.on_init:
			i.state.Init()
			i.publish("task/init", i.state)

		case req := <-i.on_complete:
			i.state.Complete(req.Result)
			i.publish("task/complete", i.state)
			return

		case req := <-i.on_fail:
			i.state.Fail(req.Error)
			i.publish("task/fail", i.state)
			return

		case req := <-i.on_log:
			log, exists := i.logs[req.File]
			if !exists {
				log = make([]string, 0, 32)
			}
			i.logs[req.File] = append(log, req.Data)

		case <-ctx.Done():
			i.state.Fail(fmt.Errorf("killed by task manager: timeout exceeded"))
			i.publish("task/fail", i.state)
			return

		case <-time.After(10 * time.Second):
			// periodic liveness check
			fmt.Println("poke", i.id)
			if err := i.driver.Poke(ctx, i.id); err != nil {
				fmt.Println("task", i.id, "failed liveness check:", err)
				i.state.Fail(fmt.Errorf("cluster task error: %w", err))
				i.publish("task/fail", i.state)
				return
			}
		}
	}
}

func (i *instance) cleanup() {
	defer close(i.on_init)
	defer close(i.on_complete)
	defer close(i.on_fail)
	defer close(i.on_log)

	// wait a sec for any logs to arrive
	// todo: avoid race condition here
	time.Sleep(time.Second)

	// delete completed tasks
	if i.state.Status == task.StatusDone {
		ctx := context.Background()
		if err := i.driver.Kill(ctx, i.id); err != nil {
			// log error
			fmt.Println("failed to kill", i, ":", err)
		}
	}
}
