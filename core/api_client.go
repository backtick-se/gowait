package core

import (
	"context"
	"github.com/backtick-se/gowait/core/task"
)

type APIClient interface {
	Connect(hostname string) error

	CreateTask(context.Context, *task.Spec) (*task.State, error)
}
