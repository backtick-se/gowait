package client

import (
	"context"
	"cowait/core"
	"cowait/core/pb"

	"google.golang.org/grpc"
)

type Client interface {
	Connect(hostname string) error
	CreateTask(context.Context, core.TaskID, *core.TaskSpec) (*core.TaskState, error)
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

func (c *client) CreateTask(ctx context.Context, id core.TaskID, def *core.TaskSpec) (*core.TaskState, error) {
	if c.conn == nil {
		return nil, ErrNotConnected
	}
	reply, err := c.api.CreateTask(ctx, &pb.CreateTaskReq{
		Id:   string(id),
		Spec: pb.PackTaskSpec(def),
	})
	if err != nil {
		return nil, err
	}
	return &core.TaskState{
		ID:        core.TaskID(reply.Instance.Id),
		Parent:    core.TaskID(reply.Instance.Parent),
		Status:    core.TaskStatus(reply.Instance.Status),
		Spec:      pb.UnpackTaskSpec(reply.Instance.Spec),
		Scheduled: reply.Instance.Scheduled.AsTime(),
		Started:   reply.Instance.Started.AsTime(),
		Completed: reply.Instance.Completed.AsTime(),
	}, nil
}
