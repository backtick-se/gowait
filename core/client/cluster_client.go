package client

import (
	"context"
	"cowait/core"
	"net"
)

type Cluster interface {
	Connect(net.Conn) error
	Connected() bool
	Info(context.Context) (*core.ClusterInfo, error)
	Subscribe(ctx context.Context) (ClusterEventStream, error)
}

type ClusterEventStream interface {
	Next() <-chan *core.ClusterEvent
}
