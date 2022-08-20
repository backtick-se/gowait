package daemon

import (
	"cowait/core"
	"cowait/core/msg"

	"fmt"
)

type TaskServer interface {
	Init(*msg.TaskInit) error
	Complete(*msg.TaskComplete) error
	Fail(*msg.TaskFailure) error
	Log(*msg.LogEntry) error
}

type TaskManager interface {
	TaskServer

	Get(core.TaskID) (Instance, bool)
	Schedule(*core.TaskDef) (Instance, error)
}

func NewTaskManager(cluster core.Cluster) TaskManager {
	return &taskmgr{
		cluster: cluster,
		tasks:   make(map[core.TaskID]*instance),
	}
}

func newTaskServer(taskmgr TaskManager) TaskServer {
	// re-export as a TaskServer
	return taskmgr
}

type taskmgr struct {
	cluster core.Cluster
	tasks   map[core.TaskID]*instance
}

func (t *taskmgr) Get(id core.TaskID) (i Instance, ok bool) {
	i, ok = t.tasks[id]
	return
}

func (t *taskmgr) Schedule(task *core.TaskDef) (Instance, error) {
	instance := newInstance(t.cluster, task)
	t.tasks[instance.ID()] = instance
	fmt.Printf("Scheduled task: %+v\n", task)
	return instance, nil
}

func (t *taskmgr) Init(req *msg.TaskInit) error {
	fmt.Printf("Task init: %+v\n", req)
	if task, ok := t.tasks[req.Header.ID]; ok {
		task.on_init <- req
		return nil
	}
	fmt.Println("Unknown task", req.Header.ID)
	return core.ErrUnknownTask
}

func (t *taskmgr) Complete(req *msg.TaskComplete) error {
	if task, ok := t.tasks[req.Header.ID]; ok {
		fmt.Printf("Task complete: %s\n", req.Result)
		task.on_complete <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) Fail(req *msg.TaskFailure) error {
	if task, ok := t.tasks[req.Header.ID]; ok {
		fmt.Printf("Task failed: %s\n", req.Error)
		task.on_fail <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) Log(req *msg.LogEntry) error {
	if task, ok := t.tasks[req.Header.ID]; ok {
		fmt.Printf("%s [%s] %s", req.Header.ID, req.File, req.Data)
		task.on_log <- req
	}
	return core.ErrUnknownTask
}
