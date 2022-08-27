package daemon

import "github.com/backtick-se/gowait/core/task"

type TaskManager interface {
	Get(task.ID) (Instance, bool)
	Add(Instance)
}

type taskmgr struct {
	byId map[task.ID]Instance
}

func NewTaskManager() TaskManager {
	return &taskmgr{
		byId: make(map[task.ID]Instance),
	}
}

func (t *taskmgr) Get(id task.ID) (Instance, bool) {
	i, ok := t.byId[id]
	return i, ok
}

func (t *taskmgr) Add(i Instance) {
	t.byId[i.ID()] = i
}
