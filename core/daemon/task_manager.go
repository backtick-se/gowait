package daemon

import (
	"cowait/core"
	"cowait/core/msg"

	"fmt"
)

type TaskManager interface {
	Get(core.TaskID) (Instance, bool)
	Schedule(*core.TaskDef) error

	Init(*msg.TaskInit) error
	Complete(*msg.TaskComplete) error
	Fail(*msg.TaskFailure) error
	Log(*msg.LogEntry) error
}

func NewTaskManager(cluster core.Cluster) TaskManager {
	return &taskmgr{
		cluster: cluster,
		tasks:   make(map[core.TaskID]*instance),
	}
}

type taskmgr struct {
	cluster core.Cluster
	tasks   map[core.TaskID]*instance
}

func (t *taskmgr) Get(id core.TaskID) (i Instance, ok bool) {
	i, ok = t.tasks[id]
	return
}

func (t *taskmgr) Schedule(task *core.TaskDef) error {
	t.tasks[task.ID] = newInstance(t.cluster, task)
	fmt.Printf("Scheduled task: %+v\n", task)
	return nil
}

func (t *taskmgr) Init(req *msg.TaskInit) error {
	if task, ok := t.tasks[req.ID]; ok {
		fmt.Printf("Task init: %+v\n", req)
		task.on_init <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) Complete(req *msg.TaskComplete) error {
	if task, ok := t.tasks[req.ID]; ok {
		fmt.Printf("Task complete: %s\n", req.Result)
		task.on_complete <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) Fail(req *msg.TaskFailure) error {
	if task, ok := t.tasks[req.ID]; ok {
		fmt.Printf("Task failed: %s\n", req.Error)
		task.on_fail <- req
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) Log(req *msg.LogEntry) error {
	if task, ok := t.tasks[req.ID]; ok {
		task.on_log <- req
		fmt.Printf("%s [%s] %s", req.Header.ID, req.File, req.Data)
	}
	return core.ErrUnknownTask
}
