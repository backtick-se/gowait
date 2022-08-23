package executor

import (
	"context"
	"cowait/core/msg"
	"cowait/core/task"
)

type Client interface {
	Connect(hostname string, id task.ID) error

	Init(ctx context.Context) error
	Failure(ctx context.Context, taskErr error) error
	Complete(ctx context.Context, result string) error
	Log(ctx context.Context) (Logger, error)
}

// Handles commands from executors
type Handler interface {
	Init(*msg.TaskInit) error
	Complete(*msg.TaskComplete) error
	Fail(*msg.TaskFailure) error
	Log(*msg.LogEntry) error
}

type Logger interface {
	Log(file, data string) error
	Close() error
}
