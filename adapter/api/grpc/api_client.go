package grpc

import (
	"context"
	"cowait/adapter/api/grpc/pb"
	"cowait/core"
	"cowait/core/client"

	"google.golang.org/grpc"
)

type apiclient struct {
	conn *grpc.ClientConn
	api  pb.CowaitClient
}

func NewApiClient() client.API {
	return &apiclient{}
}

func (c *apiclient) Connect(hostname string) error {
	return c.dial(hostname, grpc.WithInsecure())
}

func (c *apiclient) dial(hostname string, opts ...grpc.DialOption) error {
	var err error
	c.conn, err = grpc.Dial(hostname, opts...)
	if err != nil {
		return err
	}

	c.api = pb.NewCowaitClient(c.conn)
	return nil
}

func (c *apiclient) CreateTask(ctx context.Context, def *core.TaskSpec) (*core.TaskState, error) {
	if c.conn == nil {
		return nil, client.ErrNotConnected
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
