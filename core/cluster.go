package core

import "context"

type Cluster interface {
	Name() string
	Get(context.Context, TaskID) (Task, bool)
	Create(context.Context, *TaskSpec) (Task, error)
	Destroy(context.Context, TaskID) error
}
