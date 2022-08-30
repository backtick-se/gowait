package daemon

import (
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/task"
	"github.com/backtick-se/gowait/util"
	"github.com/backtick-se/gowait/util/events"

	"context"
	"fmt"
	"time"
)

type daemon struct {
	tasks   task.Manager
	queue   task.TaskQueue
	workers Workers

	info   cluster.Info
	events events.Pub[*cluster.Event]
}

func NewDaemon(workers Workers, taskmgr task.Manager, queue task.TaskQueue) cluster.T {
	return &daemon{
		workers: workers,
		tasks:   taskmgr,
		queue:   queue,

		events: events.New[*cluster.Event](),
		info: cluster.Info{
			ID:   util.RandomString(8),
			Name: "test-daemon",
		},
	}
}

func (t *daemon) Info() cluster.Info {
	return t.info
}

func (t *daemon) Events() events.Pub[*cluster.Event] {
	return t.events
}

func (t *daemon) Get(ctx context.Context, id task.ID) (task.T, bool) {
	return t.tasks.Get(id)
}

func (t *daemon) Create(ctx context.Context, spec *task.Spec) (task.T, error) {
	// set logical time
	if spec.Time.IsZero() {
		spec.Time = time.Now()
	}

	// create a task instance
	instance := task.NewInstance(spec)

	// add to task manager
	if err := t.tasks.Add(instance); err != nil {
		return nil, err
	}

	// add to work queue
	if err := t.queue.Queue(ctx, instance); err != nil {
		return nil, err
	}

	// request an idle worker
	// todo: mitigate race condition
	//   the worker is not marked as busy.
	//   subsequent calls to request before aquire will return the same executor.
	//   fix: introduce a Reserved state with a timeout
	if _, err := t.workers.Request(ctx, spec.Image); err != nil {
		return nil, err
	}

	return instance, nil
}

func (t *daemon) Destroy(ctx context.Context, id task.ID) error {
	instance, ok := t.tasks.Get(id)
	if !ok {
		return core.ErrUnknownTask
	}

	switch instance.State().Status {
	case task.StatusWait:
		instance.Fail(task.ErrCanceled)
		return nil

	case task.StatusExec:
		// todo: introduce a different error for aborts
		instance.Fail(task.ErrCanceled)
		// kill the executor that is running the task
		return t.workers.Remove(ctx, instance.State().Executor)

	default:
		return fmt.Errorf("%w: cant destroy task %s in state %s", ErrInvalidState, id, instance.State().Status)
	}
}
