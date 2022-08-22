package core

import (
	"context"
)

type APIClient interface {
	Connect(hostname string) error

	CreateTask(context.Context, *TaskSpec) (*TaskState, error)
}
