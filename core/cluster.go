package core

import (
	"context"
)

type Cluster interface {
	Spawn(context.Context, *TaskDef) (Task, error)
	Kill(context.Context, TaskID) error
	Poke(context.Context, TaskID) error
}
