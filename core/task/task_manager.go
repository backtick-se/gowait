package task

import (
	"github.com/backtick-se/gowait/core"

	"fmt"

	"go.uber.org/zap"
)

type Manager interface {
	Handler

	Add(Instance) error
	Get(ID) (Instance, bool)
}

type taskmgr struct {
	byId map[ID]Instance
	log  *zap.Logger
}

func NewManager(log *zap.Logger) Manager {
	return &taskmgr{
		log:  log,
		byId: make(map[ID]Instance),
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
		t.log.Info("task init", zap.String("task", string(id)), zap.String("executor", string(req.Executor)))
		instance.OnInit(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) OnComplete(req *MsgComplete) error {
	id := ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		t.log.Info("task complete", zap.String("task", string(id)), zap.String("result", string(req.Result)))
		instance.OnComplete(req)
		return nil
	}
	return core.ErrUnknownTask
}

func (t *taskmgr) OnFailure(req *MsgFailure) error {
	id := ID(req.Header.ID)
	if instance, ok := t.Get(id); ok {
		t.log.Info("task failure", zap.String("task", string(id)), zap.Error(req.Error))
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
