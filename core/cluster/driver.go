package cluster

import (
	"context"
	"github.com/backtick-se/gowait/core/task"
)

type Driver interface {
	Spawn(context.Context, task.ID, string) error
	Kill(context.Context, task.ID) error
	Poke(context.Context, task.ID) error
}
