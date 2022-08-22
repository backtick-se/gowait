package core

import (
	"context"
	"cowait/util/events"
	"net"
)

type ClusterClient interface {
	Connect(net.Conn) error
	Close() error

	Info(context.Context) (*ClusterInfo, error)
	Subscribe(ctx context.Context) (ClusterEventStream, error)
}

type ClusterEventStream interface {
	Next() <-chan *ClusterEvent
}

type Cluster interface {
	Info() ClusterInfo
	Events() events.Pub[*ClusterEvent]
	Get(context.Context, TaskID) (Task, bool)
	Create(context.Context, *TaskSpec) (Task, error)
	Destroy(context.Context, TaskID) error
}

type ClusterInfo struct {
	ID   string
	Name string
}

type ClusterEvent struct {
	ID   string
	Type string
}
