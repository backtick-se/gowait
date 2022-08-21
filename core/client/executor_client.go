package client

import (
	"cowait/core"

	"context"
)

type Executor interface {
	Connect(hostname string, id core.TaskID) error

	Init(ctx context.Context) error
	Failure(ctx context.Context, taskErr error) error
	Complete(ctx context.Context, result string) error
	Log(ctx context.Context) (Logger, error)
}

type Logger interface {
	Log(file, data string) error
	Close() error
}
