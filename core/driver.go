package core

import (
	"context"
)

type Driver interface {
	Spawn(context.Context, TaskID, *TaskSpec) error
	Kill(context.Context, TaskID) error
	Poke(context.Context, TaskID) error
}
