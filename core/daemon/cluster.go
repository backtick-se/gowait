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

type daemon struct {
	info    cluster.Info
	driver  cluster.Driver
	workers Workers
	queue   Queue
	tasks   TaskManager
	events  events.Pub[*cluster.Event]
}

func NewDaemon(driver cluster.Driver) cluster.T {
	return &daemon{
		events:  events.New[*cluster.Event](),
		driver:  driver,
		workers: NewWorkers(driver),
		queue:   NewQueue(256),
		tasks:   NewTaskManager(),
		info: cluster.Info{
			ID:   util.RandomString(8),
			Name: "test-daemon",
		},
	}
}

var _ executor.Handler = &daemon{}

func registerExecutorHandler(cluster cluster.T) (executor.Handler, error) {
	// re-export as an executor server
	if srv, ok := cluster.(executor.Handler); ok {
		return srv, nil
	}
	return nil, fmt.Errorf("expected cluster implementation to also be an executor server")
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

	// queue up a task instance
	instance, err := t.queue.Push(ctx, spec)
	if err != nil {
		return nil, err
	}

	// add to list of tasks
	t.tasks.Add(instance)

	// request an idle worker
	if _, err := t.workers.Request(ctx, spec.Image); err != nil {
		return nil, err
	}

	return instance, nil
}

func (t *daemon) Destroy(ctx context.Context, id task.ID) error {
	return t.driver.Kill(ctx, id)
}

//
// Executor Server implementation
//

func (t *daemon) ExecInit(ctx context.Context, req *msg.ExecInit) error {
	id := task.ID(req.Header.ID)
	if worker, ok := t.workers.Get(id); ok {
		fmt.Println("executor init:", id)
		worker.OnInit()
	}
	return nil
}

func (t *daemon) ExecAquire(ctx context.Context, req *msg.ExecAquire) (*task.Run, error) {
	id := task.ID(req.Header.ID)
	if worker, ok := t.workers.Get(id); ok {
		// find the next suitable work item for this executor
		// this will block until a new task is available.
		// its up to the caller to abort the call if the wait is too long.
		// this greatly reduces task startup latency
		instance, err := t.queue.Pop(ctx, worker.Image())
		if err != nil {
			return nil, err
		}

		fmt.Println("executor/aquire", *instance.State())
		worker.OnAquire(instance)
		return instance.State(), nil
	}
	return nil, core.ErrUnknownTask
}

func (t *daemon) ExecStop(ctx context.Context, req *msg.ExecStop) error {
	id := task.ID(req.Header.ID)
	if worker, ok := t.workers.Get(id); ok {
		fmt.Println("executor stopped:", id)
		worker.OnStop()
		t.workers.Remove(worker)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Init(req *msg.TaskInit) error {
	id := task.ID(req.Header.ID)
	if instance, ok := t.tasks.Get(id); ok {
		fmt.Println("task/init", id, "on", req.Executor)
		instance.OnInit(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Complete(req *msg.TaskComplete) error {
	id := task.ID(req.Header.ID)
	if instance, ok := t.tasks.Get(id); ok {
		fmt.Println("task/complete", id, string(req.Result))
		instance.OnComplete(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Fail(req *msg.TaskFailure) error {
	id := task.ID(req.Header.ID)
	if instance, ok := t.tasks.Get(id); ok {
		fmt.Println("task/fail", id, req.Error)
		instance.OnFailure(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Log(req *msg.LogEntry) error {
	id := task.ID(req.Header.ID)
	if instance, ok := t.tasks.Get(id); ok {
		fmt.Printf("%s [%s] %s", req.Header.ID, req.File, req.Data)
		instance.OnLog(req)
	}
	return core.ErrUnknownTask
}
