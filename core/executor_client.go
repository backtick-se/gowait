package core

import (
	"context"
)

type ExecutorClient interface {
	Connect(hostname string, id TaskID) error

	Init(ctx context.Context) error
	Failure(ctx context.Context, taskErr error) error
	Complete(ctx context.Context, result string) error
	Log(ctx context.Context) (Logger, error)
}

type Logger interface {
	Log(file, data string) error
	Close() error
}
