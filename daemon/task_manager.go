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

func (t *taskmgr) Init(*msg.TaskInit) error {
	return nil
}

func (t *taskmgr) Status(*msg.TaskStatus) error {
	return nil
}

func (t *taskmgr) Complete(*msg.TaskComplete) error {
	return nil
}

func (t *taskmgr) Fail(*msg.TaskFailure) error {
	return nil
}

func (t *taskmgr) Log(entry *msg.LogEntry) error {
	fmt.Printf("%s [%s] %s", entry.Header.ID, entry.File, entry.Data)
	return nil
}
