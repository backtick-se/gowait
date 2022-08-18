package client

import (
	"cowait/core"
	"cowait/core/pb"

	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client interface {
	Init(ctx context.Context) error
	Failure(ctx context.Context, taskErr error) error
	Complete(ctx context.Context, result string) error
	Log(ctx context.Context) (TaskLogger, error)
}

type client struct {
	taskID  core.TaskID
	conn    *grpc.ClientConn
	tasks   pb.TaskClient
	version string
}

func New(taskID core.TaskID) (Client, error) {
	conn, err := grpc.Dial("cowaitd.default.svc.cluster.local:1337", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	tasks := pb.NewTaskClient(conn)

	return &client{
		taskID:  taskID,
		conn:    conn,
		tasks:   tasks,
		version: "gowait/1.0",
	}, nil
}

func (c *client) header() *pb.Header {
	return &pb.Header{
		TaskId: string(c.taskID),
		Time:   timestamppb.New(time.Now()),
	}
}

func (c *client) Init(ctx context.Context) error {
	_, err := c.tasks.TaskInit(ctx, &pb.TaskInitReq{
		Header:  c.header(),
		Version: c.version,
	})
	return err
}

func (c *client) Failure(ctx context.Context, taskErr error) error {
	_, err := c.tasks.TaskFailure(ctx, &pb.TaskFailureReq{
		Header: c.header(),
		Error:  taskErr.Error(),
	})
	return err
}

func (c *client) Complete(ctx context.Context, result string) error {
	_, err := c.tasks.TaskComplete(ctx, &pb.TaskCompleteReq{
		Header: c.header(),
		Result: result,
	})
	return err
}

func (c *client) Log(ctx context.Context) (TaskLogger, error) {
	stream, err := c.tasks.TaskLog(ctx)
	if err != nil {
		return nil, err
	}
	return &taskLog{
		client: c,
		stream: stream,
	}, nil
}
