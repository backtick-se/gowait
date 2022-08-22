package daemon

import (
	"cowait/core"
	"cowait/core/msg"

	"context"
	"fmt"
	"time"
)

type task struct {
	id     core.TaskID
	driver core.Driver
	state  core.TaskState
	logs   map[string][]string

	publish     TaskEventFn
	on_init     chan *msg.TaskInit
	on_fail     chan *msg.TaskFailure
	on_complete chan *msg.TaskComplete
	on_log      chan *msg.LogEntry
}

var _ core.Task = &task{}

type TaskEventFn func(event string, state core.TaskState)

func newTask(driver core.Driver, id core.TaskID, spec *core.TaskSpec, callback TaskEventFn) *task {
	t := &task{
		id:     id,
		driver: driver,
		state: core.TaskState{
			ID:        id,
			Spec:      spec,
			Status:    core.StatusWait,
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

func (i *task) ID() core.TaskID       { return i.state.ID }
func (i *task) Spec() *core.TaskSpec  { return i.state.Spec }
func (i *task) State() core.TaskState { return i.state }

func (i *task) Logs(file string) []string {
	if logs, ok := i.logs[file]; ok {
		return logs
	}
	return nil
}

func (i *task) proc() {
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

func (i *task) cleanup() {
	defer close(i.on_init)
	defer close(i.on_complete)
	defer close(i.on_fail)
	defer close(i.on_log)

	// wait a sec for any logs to arrive
	// todo: avoid race condition here
	time.Sleep(time.Second)

	// delete completed tasks
	if i.state.Status == core.StatusDone {
		ctx := context.Background()
		if err := i.driver.Kill(ctx, i.id); err != nil {
			// log error
			fmt.Println("failed to kill", i, ":", err)
		}
	}
}
