package daemon

import (
	"context"
	"cowait/core"
	"cowait/core/msg"
	"cowait/util/events"

	"fmt"
	"time"
)

func NewCluster(driver core.Driver) core.Cluster {
	return &cluster{
		events: events.New[*core.ClusterEvent](),
		driver: driver,
		tasks:  make(map[core.TaskID]*task),
		info: core.ClusterInfo{
			ID:   "4123",
			Name: "test-daemon",
		},
	}
}

func registerExecutorHandler(cluster core.Cluster) (core.ExecutorHandler, error) {
	// re-export as an executor server
	if srv, ok := cluster.(core.ExecutorHandler); ok {
		return srv, nil
	}
	return nil, fmt.Errorf("expected cluster implementation to also be an executor server")
}

type cluster struct {
	info   core.ClusterInfo
	driver core.Driver
	tasks  map[core.TaskID]*task
	events events.Pub[*core.ClusterEvent]
}

func (t *cluster) Info() core.ClusterInfo {
	return t.info
}

func (t *cluster) Events() events.Pub[*core.ClusterEvent] {
	return t.events
}

func (t *cluster) Get(ctx context.Context, id core.TaskID) (i core.Task, ok bool) {
	i, ok = t.tasks[id]
	return
}

func (t *cluster) Create(ctx context.Context, task *core.TaskSpec) (core.Task, error) {
	// generate task id
	id := core.GenerateTaskID(task.Name)

	// set logical time
	if task.Time.IsZero() {
		task.Time = time.Now()
	}

	fmt.Printf("Create task %s: %+v\n", id, task)
	instance := newTask(t.driver, id, task)
	t.tasks[id] = instance

	t.events.Publish(&core.ClusterEvent{
		ID:   t.info.ID,
		Type: "task/create",
	})

	return instance, nil
}

func (t *cluster) Destroy(ctx context.Context, id core.TaskID) error {
	return t.driver.Kill(ctx, id)
}

//
// Executor Server implementation
//

func (t *cluster) Init(req *msg.TaskInit) error {
	id := core.TaskID(req.Header.ID)
	fmt.Printf("Task init: %+v\n", req)
	if task, ok := t.tasks[id]; ok {
		task.on_init <- req
		return nil
	}
	fmt.Println("Unknown task", req.Header.ID)
	return core.ErrUnknownTask
}

func (t *cluster) Complete(req *msg.TaskComplete) error {
	id := core.TaskID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Printf("Task complete: %s\n", req.Result)
		task.on_complete <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *cluster) Fail(req *msg.TaskFailure) error {
	id := core.TaskID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Printf("Task failed: %s\n", req.Error)
		task.on_fail <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *cluster) Log(req *msg.LogEntry) error {
	id := core.TaskID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Printf("%s [%s] %s", req.Header.ID, req.File, req.Data)
		task.on_log <- req
	}
	return core.ErrUnknownTask
}
