package client

import (
	"context"
	"net"
)

type Cluster interface {
	Connect(net.Conn) error
	Connected() bool
	Info(context.Context) (*ClusterInfo, error)
	Subscribe(ctx context.Context) (ClusterEventStream, error)
}

type ClusterEventStream interface {
	Read() (*ClusterEvent, bool)
}

type ClusterInfo struct {
	Name string
}

type ClusterEvent struct {
	Type string
}
