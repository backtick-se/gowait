package daemon

import (
	"context"
	"fmt"

	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/msg"
	"github.com/backtick-se/gowait/core/task"
)

// todo: move to task package
type TaskHandler interface {
	OnInit(*msg.TaskInit) error
	OnComplete(*msg.TaskComplete) error
	OnFailure(*msg.TaskFailure) error
	OnLog(*msg.LogEntry) error
}

type TaskManager interface {
	TaskQueue
	TaskHandler

	Get(task.ID) (Instance, bool)
}

type taskmgr struct {
	TaskQueue
	byId map[task.ID]Instance
}

func NewTaskManager() TaskManager {
	return &taskmgr{
		TaskQueue: NewTaskQueue(256),
		byId:      make(map[task.ID]Instance),
	}
}

func (t *taskmgr) Get(id task.ID) (Instance, bool) {
	i, ok := t.byId[id]
	return i, ok
}

func (t *taskmgr) Queue(ctx context.Context, spec *task.Spec) (Instance, error) {
	instance, err := t.TaskQueue.Queue(ctx, spec)
	if err != nil {
		return nil, err
	}
	t.byId[instance.ID()] = instance
	return instance, nil
}

//
// task.Handler implementation
//

func (t *taskmgr) OnInit(req *msg.TaskInit) error {
	id := task.ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		fmt.Println("task/init", id, "on", req.Executor)
		instance.OnInit(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) OnComplete(req *msg.TaskComplete) error {
	id := task.ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		fmt.Println("task/complete", id, string(req.Result))
		instance.OnComplete(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) OnFailure(req *msg.TaskFailure) error {
	id := task.ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		fmt.Println("task/fail", id, req.Error)
		instance.OnFailure(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) OnLog(req *msg.LogEntry) error {
	id := task.ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		fmt.Printf("%s [%s] %s", req.Header.ID, req.File, req.Data)
		instance.OnLog(req)
	}
	return core.ErrUnknownTask
}
