package daemon

import (
	"cowait/core"
	"cowait/core/msg"
	"time"

	"fmt"
)

type TaskManager interface {
	ExecutorServer
	ClusterServer
}

func NewTaskManager(cluster core.Cluster) TaskManager {
	return &taskmgr{
		cluster: cluster,
		tasks:   make(map[core.TaskID]*instance),
	}
}

func newTaskServer(taskmgr TaskManager) ExecutorServer {
	// re-export as an executor server
	return taskmgr
}

type taskmgr struct {
	cluster core.Cluster
	tasks   map[core.TaskID]*instance
}

func (t *taskmgr) Get(id core.TaskID) (i core.Task, ok bool) {
	i, ok = t.tasks[id]
	return
}

func (t *taskmgr) Schedule(task *core.TaskSpec) (core.Task, error) {
	// generate task id
	id := core.GenerateTaskID(task.Name)

	// set logical time
	if task.Time.IsZero() {
		task.Time = time.Now()
	}

	instance := newInstance(t.cluster, id, task)
	t.tasks[id] = instance
	fmt.Printf("Scheduled task %s: %+v\n", id, task)
	return instance, nil
}

func (t *taskmgr) Init(req *msg.TaskInit) error {
	id := core.TaskID(req.Header.ID)
	fmt.Printf("Task init: %+v\n", req)
	if task, ok := t.tasks[id]; ok {
		task.on_init <- req
		return nil
	}
	fmt.Println("Unknown task", req.Header.ID)
	return core.ErrUnknownTask
}

func (t *taskmgr) Complete(req *msg.TaskComplete) error {
	id := core.TaskID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Printf("Task complete: %s\n", req.Result)
		task.on_complete <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) Fail(req *msg.TaskFailure) error {
	id := core.TaskID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Printf("Task failed: %s\n", req.Error)
		task.on_fail <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) Log(req *msg.LogEntry) error {
	id := core.TaskID(req.Header.ID)
	if task, ok := t.tasks[id]; ok {
		fmt.Printf("%s [%s] %s", req.Header.ID, req.File, req.Data)
		task.on_log <- req
	}
	return core.ErrUnknownTask
}
