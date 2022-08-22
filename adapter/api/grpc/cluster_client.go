package grpc

import (
	"cowait/adapter/api/grpc/pb"
	"cowait/core/client"
	"fmt"
	"io"

	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type clusterClient struct {
	conn    *grpc.ClientConn
	cluster pb.ClusterClient
}

func NewClusterClient() client.Cluster {
	return &clusterClient{}
}

func (c *clusterClient) Connect(conn net.Conn) error {
	dialer := func(context.Context, string) (net.Conn, error) {
		return conn, nil
	}

	var err error
	c.conn, err = grpc.Dial(":1", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
	if err != nil {
		return err
	}

	c.cluster = pb.NewClusterClient(c.conn)
	return nil
}

func (c *clusterClient) Connected() bool {
	return c.conn.GetState() == connectivity.Ready
}

func (c *clusterClient) Info(ctx context.Context) (*client.ClusterInfo, error) {
	reply, err := c.cluster.Info(ctx, &pb.ClusterInfoReq{})
	if err != nil {
		return nil, err
	}
	return &client.ClusterInfo{
		Name: reply.Name,
	}, nil
}

func (c *clusterClient) Subscribe(ctx context.Context) (client.ClusterEventStream, error) {
	stream, err := c.cluster.Subscribe(ctx, &pb.ClusterSubscribeReq{})
	if err != nil {
		return nil, err
	}
	events := &clusterEventStream{
		stream: stream,
	}
	go events.proc()
	return events, nil
}

type clusterEventStream struct {
	stream pb.Cluster_SubscribeClient
	events chan *client.ClusterEvent
}

func (c *clusterEventStream) proc() {
	defer close(c.events)
	for {
		event, err := c.stream.Recv()
		if err != nil {
			if err != io.EOF {
				fmt.Println("cluster event stream failure:", err)
			}
			return
		}
		c.events <- &client.ClusterEvent{
			Type: event.Type,
		}
	}
}

func (c *clusterEventStream) Read() (*client.ClusterEvent, bool) {
	event, ok := <-c.events
	return event, ok
}
