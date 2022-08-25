package daemon

import (
	"context"
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/msg"
	"github.com/backtick-se/gowait/core/task"
	"github.com/backtick-se/gowait/util"
	"github.com/backtick-se/gowait/util/events"

	"fmt"
	"time"
)

func NewDaemon(driver cluster.Driver) cluster.T {
	return &daemon{
		events: events.New[*cluster.Event](),
		driver: driver,
		tasks:  make(map[task.ID]*instance),
		info: cluster.Info{
			ID:   util.RandomString(8),
			Name: "test-daemon",
		},
	}
}

func registerExecutorHandler(cluster cluster.T) (executor.Handler, error) {
	// re-export as an executor server
	if srv, ok := cluster.(executor.Handler); ok {
		return srv, nil
	}
	return nil, fmt.Errorf("expected cluster implementation to also be an executor server")
}

type daemon struct {
	info   cluster.Info
	driver cluster.Driver
	tasks  map[task.ID]*instance
	events events.Pub[*cluster.Event]
}

func (t *daemon) Info() cluster.Info {
	return t.info
}

func (t *daemon) Events() events.Pub[*cluster.Event] {
	return t.events
}

func (t *daemon) Get(ctx context.Context, id task.ID) (i task.T, ok bool) {
	i, ok = t.tasks[id]
	return
}

func (t *daemon) publish(event string, state task.State) {
	fmt.Println(event, state.ID)
	t.events.Publish(&cluster.Event{
		ID:   t.info.ID,
		Type: event,
		Task: state,
	})
}

func (t *daemon) Create(ctx context.Context, tsk *task.Spec) (task.T, error) {
	// generate task id
	id := task.GenerateID(tsk.Name)

	// set logical time
	if tsk.Time.IsZero() {
		tsk.Time = time.Now()
	}

	instance := newInstance(t.driver, id, tsk, t.publish)
	t.tasks[id] = instance

	return instance, nil
}

func (t *daemon) Destroy(ctx context.Context, id task.ID) error {
	return t.driver.Kill(ctx, id)
}

//
// Executor Server implementation
//

func (t *daemon) Init(req *msg.TaskInit) error {
	id := task.ID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		task.on_init <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Complete(req *msg.TaskComplete) error {
	id := task.ID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		task.on_complete <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Fail(req *msg.TaskFailure) error {
	id := task.ID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		task.on_fail <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Log(req *msg.LogEntry) error {
	id := task.ID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Printf("%s [%s] %s", req.Header.ID, req.File, req.Data)
		task.on_log <- req
	}
	return core.ErrUnknownTask
}
