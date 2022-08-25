package grpc

import (
	"github.com/backtick-se/gowait/adapter/api/grpc/pb"
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type executorClient struct {
	taskID  task.ID
	conn    *grpc.ClientConn
	tasks   pb.ExecutorClient
	version string
}

func NewExecutorClient() executor.Client {
	return &executorClient{
		version: "gowait/1.0",
	}
}

func (c *executorClient) Connect(hostname string, id task.ID) error {
	var err error
	c.conn, err = grpc.Dial(hostname, grpc.WithInsecure())
	if err != nil {
		return err
	}

	c.taskID = id
	c.tasks = pb.NewExecutorClient(c.conn)
	return nil
}

func (c *executorClient) header() *pb.Header {
	return &pb.Header{
		Id:   string(c.taskID),
		Time: timestamppb.New(time.Now()),
	}
}

func (c *executorClient) Init(ctx context.Context) error {
	if c.conn == nil {
		return core.ErrNotConnected
	}
	_, err := c.tasks.TaskInit(ctx, &pb.TaskInitReq{
		Header:  c.header(),
		Version: c.version,
	})
	return err
}

func (c *executorClient) Failure(ctx context.Context, taskErr error) error {
	if c.conn == nil {
		return core.ErrNotConnected
	}
	_, err := c.tasks.TaskFailure(ctx, &pb.TaskFailureReq{
		Header: c.header(),
		Error:  taskErr.Error(),
	})
	return err
}

func (c *executorClient) Complete(ctx context.Context, result string) error {
	if c.conn == nil {
		return core.ErrNotConnected
	}
	_, err := c.tasks.TaskComplete(ctx, &pb.TaskCompleteReq{
		Header: c.header(),
		Result: result,
	})
	return err
}

func (c *executorClient) Log(ctx context.Context) (executor.Logger, error) {
	if c.conn == nil {
		return nil, core.ErrNotConnected
	}
	stream, err := c.tasks.TaskLog(ctx)
	if err != nil {
		return nil, err
	}
	return &execLog{
		client: c,
		stream: stream,
	}, nil
}
