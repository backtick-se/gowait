package client

import (
	"context"
	"cowait/core"
	"cowait/core/pb"

	"google.golang.org/grpc"
)

type Client interface {
	Connect(hostname string) error
	CreateTask(context.Context, *core.TaskDef) (core.TaskID, error)
}

type client struct {
	conn *grpc.ClientConn
	api  pb.CowaitClient
}

func NewCowaitClient() Client {
	return &client{}
}

func (c *client) Connect(hostname string) error {
	var err error
	c.conn, err = grpc.Dial(hostname, grpc.WithInsecure())
	if err != nil {
		return err
	}

	c.api = pb.NewCowaitClient(c.conn)
	return nil
}

func (c *client) CreateTask(ctx context.Context, def *core.TaskDef) (core.TaskID, error) {
	var empty core.TaskID
	if c.conn == nil {
		return empty, ErrNotConnected
	}
	reply, err := c.api.CreateTask(ctx, &pb.CreateTaskReq{
		Task: pb.PackTaskdef(def),
	})
	if err != nil {
		return empty, err
	}
	return core.TaskID(reply.Instance.Id), nil
}
