package client

import (
	"context"
	"cowait/core"
)

type API interface {
	Connect(hostname string) error

	CreateTask(context.Context, *core.TaskSpec) (*core.TaskState, error)
}
