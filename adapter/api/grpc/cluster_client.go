package grpc

import (
	"cowait/adapter/api/grpc/pb"
	"cowait/core/cluster"

	"context"
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"
)

type clusterClient struct {
	conn    *grpc.ClientConn
	cluster pb.ClusterClient
}

func NewClusterClient() cluster.Client {
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

func (c *clusterClient) Close() error {
	return c.conn.Close()
}

func (c *clusterClient) Info(ctx context.Context) (*cluster.Info, error) {
	reply, err := c.cluster.Info(ctx, &pb.ClusterInfoReq{})
	if err != nil {
		return nil, err
	}
	return &cluster.Info{
		Name: reply.Name,
	}, nil
}

func (c *clusterClient) Subscribe(ctx context.Context) (cluster.EventStream, error) {
	stream, err := c.cluster.Subscribe(ctx, &pb.ClusterSubscribeReq{})
	if err != nil {
		return nil, err
	}
	events := &clusterEventStream{
		stream: stream,
		events: make(chan *cluster.Event),
	}
	go events.proc()
	return events, nil
}

type clusterEventStream struct {
	stream pb.Cluster_SubscribeClient
	events chan *cluster.Event
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
		c.events <- &cluster.Event{
			Type: event.Type,
			Task: pb.UnpackTaskState(event.Task),
		}
	}
}

func (c *clusterEventStream) Next() <-chan *cluster.Event {
	return c.events
}
