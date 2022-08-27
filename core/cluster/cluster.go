package cluster

import (
	"github.com/backtick-se/gowait/core/task"
	"github.com/backtick-se/gowait/util/events"

	"context"
	"net"
)

type Client interface {
	Connect(net.Conn) error
	Close() error

	Info(context.Context) (*Info, error)
	Subscribe(ctx context.Context) (EventStream, error)

	CreateTask(context.Context, *task.Spec) (*task.Run, error)
}

type EventStream interface {
	Next() <-chan *Event
}

type T interface {
	Info() Info
	Events() events.Pub[*Event]

	Get(context.Context, task.ID) (task.T, bool)
	Create(context.Context, *task.Spec) (task.T, error)
	Destroy(context.Context, task.ID) error
}

type Info struct {
	ID   string
	Name string
	Key  string
}

type Event struct {
	ID   string
	Type string
	Task task.Run
}
