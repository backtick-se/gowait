package core

import (
	"context"
)

type Cluster interface {
	Spawn(context.Context, *TaskDef) (Task, error)
}
