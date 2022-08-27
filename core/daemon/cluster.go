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
	info      cluster.Info
	driver    cluster.Driver
	tasks     map[task.ID]*instance
	executors map[task.ID]*worker
	events    events.Pub[*cluster.Event]
	queue     *Queue
}

func NewDaemon(driver cluster.Driver) cluster.T {
	return &daemon{
		events:    events.New[*cluster.Event](),
		driver:    driver,
		tasks:     make(map[task.ID]*instance),
		executors: make(map[task.ID]*worker),
		queue:     NewQueue(),
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

func (t *daemon) Get(ctx context.Context, id task.ID) (i task.T, ok bool) {
	i, ok = t.tasks[id]
	return
}

func (t *daemon) publish(event string, state task.Run) {
	fmt.Println("!", event, state)
	t.events.Publish(&cluster.Event{
		ID:   t.info.ID,
		Type: event,
		Task: state,
	})
}

func (t *daemon) Create(ctx context.Context, spec *task.Spec) (task.T, error) {
	// generate task id
	id := task.GenerateID("executor")

	// set logical time
	if spec.Time.IsZero() {
		spec.Time = time.Now()
	}

	worker := newWorker(t.driver, id, spec.Image, t.publish)
	t.executors[id] = worker

	instance := newInstance(spec)
	t.queue.Enqueue(instance)
	t.queue.Enqueue(newInstance(spec))

	return instance, nil
}

func (t *daemon) Destroy(ctx context.Context, id task.ID) error {
	return t.driver.Kill(ctx, id)
}

//
// Executor Server implementation
//

func (t *daemon) ExecInit(req *msg.ExecInit) error {
	fmt.Println("executor init from", req.Header.ID)
	return nil
}

func (t *daemon) ExecAquire(req *msg.ExecAquire) (*task.Run, error) {
	id := task.ID(req.Header.ID)
	if worker, ok := t.executors[id]; ok {
		instance := t.queue.Dequeue(worker.image)
		if instance != nil {
			t.publish("executor/aquire", *instance.run)
			t.tasks[instance.ID()] = instance
			worker.on_dequeue <- instance
			return instance.run, nil
		}
		t.publish("executor/empty", task.Run{})
	}
	return nil, core.ErrUnknownTask
}

func (t *daemon) ExecStop(req *msg.ExecStop) error {
	fmt.Println("executor stopped:", req.Header.ID)
	delete(t.executors, task.ID(req.Header.ID))
	return nil
}

func (t *daemon) Init(req *msg.TaskInit) error {
	id := task.ID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Println("task/init", id, "on", req.Executor)
		task.on_init <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Complete(req *msg.TaskComplete) error {
	id := task.ID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Println("task/complete", id, string(req.Result))
		task.on_complete <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *daemon) Fail(req *msg.TaskFailure) error {
	id := task.ID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Println("task/fail", id, req.Error)
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
