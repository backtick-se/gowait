package task

import (
	"fmt"

	"github.com/backtick-se/gowait/core"
)

type Manager interface {
	Handler

	Add(Instance) error
	Get(ID) (Instance, bool)
}

type taskmgr struct {
	TaskQueue
	byId map[ID]Instance
}

func NewManager() Manager {
	return &taskmgr{
		TaskQueue: NewTaskQueue(256),
		byId:      make(map[ID]Instance),
	}
}

func (t *taskmgr) Get(id ID) (Instance, bool) {
	i, ok := t.byId[id]
	return i, ok
}

func (t *taskmgr) Add(instance Instance) error {
	t.byId[instance.ID()] = instance
	return nil
}

//
// task.Handler implementation
//

func (t *taskmgr) OnInit(req *MsgInit) error {
	id := ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		fmt.Println("task/init", id, "on", req.Executor)
		instance.OnInit(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) OnComplete(req *MsgComplete) error {
	id := ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		fmt.Println("task/complete", id, string(req.Result))
		instance.OnComplete(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) OnFailure(req *MsgFailure) error {
	id := ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		fmt.Println("task/fail", id, req.Error)
		instance.OnFailure(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) OnLog(req *MsgLog) error {
	id := ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		fmt.Printf("%s [%s] %s", req.Header.ID, req.File, req.Data)
		instance.OnLog(req)
	}
	return core.ErrUnknownTask
}
