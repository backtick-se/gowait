package daemon

import (
	"cowait/core"
	"cowait/core/msg"

	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Instance interface {
	Task() *core.TaskDef
	Status() core.TaskStatus
	Scheduled() time.Time
	Started() time.Time
	Completed() time.Time
	Result() json.RawMessage
	Err() error
	Logs(file string) []string
}

type instance struct {
	cluster   core.Cluster
	taskdef   *core.TaskDef
	status    core.TaskStatus
	scheduled time.Time
	started   time.Time
	completed time.Time
	result    json.RawMessage
	err       error
	logs      map[string][]string

	on_init     chan *msg.TaskInit
	on_fail     chan *msg.TaskFailure
	on_complete chan *msg.TaskComplete
	on_log      chan *msg.LogEntry
}

func newInstance(cluster core.Cluster, taskdef *core.TaskDef) *instance {
	i := &instance{
		cluster:   cluster,
		taskdef:   taskdef,
		status:    core.StatusWait,
		scheduled: time.Now(),
		logs:      make(map[string][]string),

		on_init:     make(chan *msg.TaskInit),
		on_fail:     make(chan *msg.TaskFailure),
		on_complete: make(chan *msg.TaskComplete),
		on_log:      make(chan *msg.LogEntry),
	}
	go i.proc()
	return i
}

func (i *instance) Task() *core.TaskDef     { return i.taskdef }
func (i *instance) Status() core.TaskStatus { return i.status }
func (i *instance) Scheduled() time.Time    { return i.scheduled }
func (i *instance) Started() time.Time      { return i.started }
func (i *instance) Completed() time.Time    { return i.completed }
func (i *instance) Result() json.RawMessage { return i.result }
func (i *instance) Err() error              { return i.err }

func (i *instance) Logs(file string) []string {
	if logs, ok := i.logs[file]; ok {
		return logs
	}
	return nil
}

func (i *instance) proc() {
	defer i.cleanup()

	// a specific context for each task allows us to set per-task deadlines etc
	ctx := context.Background()

	if i.taskdef.Timeout > 0 {
		deadline, cancel := context.WithTimeout(ctx, time.Duration(i.taskdef.Timeout)*time.Second)
		defer cancel()
		ctx = deadline
	}

	// this is the instance management loop
	// at this point the task is in the "scheduled" state
	// i suppose we start by calling cluster.Spawn() ?
	if err := i.cluster.Spawn(ctx, i.taskdef); err != nil {
		fmt.Println("failed to spawn task", i.taskdef.ID, ":", err)
		return
	}

	// todo: this should be structured as a finite state machine

	for {
		select {
		case <-i.on_init:
			i.init()

		case req := <-i.on_complete:
			i.complete(req.Result)
			return

		case req := <-i.on_fail:
			i.fail(req.Error)
			return

		case req := <-i.on_log:
			log, exists := i.logs[req.File]
			if !exists {
				log = make([]string, 0, 32)
			}
			i.logs[req.File] = append(log, req.Data)

		case <-ctx.Done():
			i.fail(fmt.Errorf("killed by task manager: timeout exceeded"))
			return

		case <-time.After(10 * time.Second):
			// periodic liveness check
			fmt.Println("poke", i.taskdef.ID)
			if err := i.cluster.Poke(ctx, i.taskdef.ID); err != nil {
				fmt.Println("task", i.taskdef.ID, "failed liveness check:", err)
				i.fail(fmt.Errorf("cluster task error: %w", err))
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

	// ensure task is gone
	ctx := context.Background()
	if err := i.cluster.Kill(ctx, i.taskdef.ID); err != nil {
		// log error
		fmt.Println("failed to kill", i, ":", err)
	}
}

func (i *instance) init() {
	i.status = core.StatusExec
	i.started = time.Now()
}

func (i *instance) complete(result json.RawMessage) {
	i.result = result
	i.status = core.StatusDone
	i.completed = time.Now()
}

func (i *instance) fail(err error) {
	i.err = err
	i.status = core.StatusFail
	i.completed = time.Now()
}
