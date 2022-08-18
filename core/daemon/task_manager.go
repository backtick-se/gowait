package daemon

import (
	"cowait/core/msg"
	"fmt"
)

type TaskManager interface {
	Init(*msg.TaskInit) error
	Status(*msg.TaskStatus) error
	Complete(*msg.TaskComplete) error
	Fail(*msg.TaskFailure) error
	Log(*msg.LogEntry) error
}

func NewTaskManager() TaskManager {
	return &taskmgr{}
}

type taskmgr struct{}

func (t *taskmgr) Init(req *msg.TaskInit) error {
	fmt.Printf("Task init: %+v\n", req)
	return nil
}

func (t *taskmgr) Status(*msg.TaskStatus) error {
	return nil
}

func (t *taskmgr) Complete(req *msg.TaskComplete) error {
	fmt.Printf("Task complete: %s\n", req.Result)
	return nil
}

func (t *taskmgr) Fail(*msg.TaskFailure) error {
	return nil
}

func (t *taskmgr) Log(entry *msg.LogEntry) error {
	fmt.Printf("%s [%s] %s", entry.Header.ID, entry.File, entry.Data)
	return nil
}
