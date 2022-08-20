package core

import (
	"context"
)

type Cluster interface {
	Name() string
	Spawn(context.Context, TaskID, *TaskSpec) error
	Kill(context.Context, TaskID) error
	Poke(context.Context, TaskID) error
}
