package client

import (
	"cowait/core"
	"cowait/core/pb"

	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskClient interface {
	Connect(hostname string, id core.TaskID) error
	Init(ctx context.Context) error
	Failure(ctx context.Context, taskErr error) error
	Complete(ctx context.Context, result string) error
	Log(ctx context.Context) (TaskLogger, error)
}

type taskClient struct {
	taskID  core.TaskID
	conn    *grpc.ClientConn
	tasks   pb.TaskClient
	version string
}

func NewTaskClient() TaskClient {
	return &taskClient{
		version: "gowait/1.0",
	}
}

func (c *taskClient) Connect(hostname string, id core.TaskID) error {
	var err error
	c.conn, err = grpc.Dial(hostname, grpc.WithInsecure())
	if err != nil {
		return err
	}

	c.taskID = id
	c.tasks = pb.NewTaskClient(c.conn)
	return nil
}

func (c *taskClient) header() *pb.Header {
	return &pb.Header{
		Id:   string(c.taskID),
		Time: timestamppb.New(time.Now()),
	}
}

func (c *taskClient) Init(ctx context.Context) error {
	if c.conn == nil {
		return ErrNotConnected
	}
	_, err := c.tasks.TaskInit(ctx, &pb.TaskInitReq{
		Header:  c.header(),
		Version: c.version,
	})
	return err
}

func (c *taskClient) Failure(ctx context.Context, taskErr error) error {
	if c.conn == nil {
		return ErrNotConnected
	}
	_, err := c.tasks.TaskFailure(ctx, &pb.TaskFailureReq{
		Header: c.header(),
		Error:  taskErr.Error(),
	})
	return err
}

func (c *taskClient) Complete(ctx context.Context, result string) error {
	if c.conn == nil {
		return ErrNotConnected
	}
	_, err := c.tasks.TaskComplete(ctx, &pb.TaskCompleteReq{
		Header: c.header(),
		Result: result,
	})
	return err
}

func (c *taskClient) Log(ctx context.Context) (TaskLogger, error) {
	if c.conn == nil {
		return nil, ErrNotConnected
	}
	stream, err := c.tasks.TaskLog(ctx)
	if err != nil {
		return nil, err
	}
	return &taskLog{
		client: c,
		stream: stream,
	}, nil
}
