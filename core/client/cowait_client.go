package client

import (
	"context"
	"cowait/core"
	"cowait/core/pb"

	"google.golang.org/grpc"
)

type Client interface {
	Connect(hostname string) error
	CreateTask(context.Context, *core.TaskSpec) (*core.TaskState, error)
}

type client struct {
	conn *grpc.ClientConn
	api  pb.CowaitClient
}

func NewCowaitClient() Client {
	return &client{}
}

func (c *client) Connect(hostname string) error {
	return c.dial(hostname, grpc.WithInsecure())
}

func (c *client) dial(hostname string, opts ...grpc.DialOption) error {
	var err error
	c.conn, err = grpc.Dial(hostname, opts...)
	if err != nil {
		return err
	}

	c.api = pb.NewCowaitClient(c.conn)
	return nil
}

func (c *client) CreateTask(ctx context.Context, def *core.TaskSpec) (*core.TaskState, error) {
	if c.conn == nil {
		return nil, ErrNotConnected
	}
	reply, err := c.api.CreateTask(ctx, &pb.CreateTaskReq{
		Spec: pb.PackTaskSpec(def),
	})
	if err != nil {
		return nil, err
	}
	return &core.TaskState{
		ID:        core.TaskID(reply.Task.TaskId),
		Parent:    core.TaskID(reply.Task.Parent),
		Status:    core.TaskStatus(reply.Task.Status),
		Spec:      pb.UnpackTaskSpec(reply.Task.Spec),
		Scheduled: reply.Task.Scheduled.AsTime(),
		Started:   reply.Task.Started.AsTime(),
		Completed: reply.Task.Completed.AsTime(),
		Result:    []byte{},
		Err:       err,
	}, nil
}
