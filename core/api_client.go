package core

import (
	"context"
	"cowait/core/task"
)

type APIClient interface {
	Connect(hostname string) error

	CreateTask(context.Context, *task.Spec) (*task.State, error)
}
