package executor

import (
	"github.com/backtick-se/gowait/core/task"

	"context"
)

type Client interface {
	Connect(hostname string, id task.ID) error
	ExecInit(context.Context, []*task.Spec) error
	ExecAquire(ctx context.Context) (*task.Run, error)
	ExecStop(ctx context.Context) error

	Init(ctx context.Context, id task.ID) error
	Failure(ctx context.Context, id task.ID, taskErr error) error
	Complete(ctx context.Context, id task.ID, result string) error
	Log(ctx context.Context, id task.ID) (Logger, error)
}

// Handles commands from executors
type Handler interface {
	ExecInit(context.Context, *MsgInit) error
	ExecAquire(context.Context, *MsgAquire) (*task.Run, error)
	ExecStop(context.Context, *MsgStop) error
}

type Logger interface {
	Log(file, data string) error
	Close() error
}
