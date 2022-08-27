package grpc

import (
	"context"
	"github.com/backtick-se/gowait/adapter/api/grpc/pb"
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/task"

	"google.golang.org/grpc"
)

type apiclient struct {
	conn *grpc.ClientConn
	api  pb.CowaitClient
}

func NewApiClient() core.APIClient {
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

func (c *apiclient) CreateTask(ctx context.Context, def *task.Spec) (*task.Run, error) {
	if c.conn == nil {
		return nil, core.ErrNotConnected
	}
	reply, err := c.api.CreateTask(ctx, &pb.CreateTaskReq{
		Spec: pb.PackTaskSpec(def),
	})
	if err != nil {
		return nil, err
	}
	return &task.Run{
		ID:        task.ID(reply.Task.TaskId),
		Parent:    task.ID(reply.Task.Parent),
		Status:    task.Status(reply.Task.Status),
		Spec:      pb.UnpackTaskSpec(reply.Task.Spec),
		Scheduled: reply.Task.Scheduled.AsTime(),
		Started:   reply.Task.Started.AsTime(),
		Completed: reply.Task.Completed.AsTime(),
		Result:    []byte{},
		Err:       err,
	}, nil
}
