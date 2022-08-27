package grpc

import (
	"github.com/backtick-se/gowait/adapter/api/grpc/pb"
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"
	"github.com/backtick-se/gowait/util/slices"

	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type executorClient struct {
	id      task.ID
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

	c.id = id
	c.tasks = pb.NewExecutorClient(c.conn)
	return nil
}

func (c *executorClient) header(id string) *pb.Header {
	return &pb.Header{
		Id:   id,
		Time: timestamppb.New(time.Now()),
	}
}

func (c *executorClient) ExecInit(ctx context.Context, specs []*task.Spec) error {
	if c.conn == nil {
		return core.ErrNotConnected
	}
	_, err := c.tasks.ExecInit(ctx, &pb.ExecInitReq{
		Header:  c.header(string(c.id)),
		Version: c.version,
		Specs:   slices.Map(specs, pb.PackTaskSpec),
	})
	return err
}

func (c *executorClient) ExecAquire(ctx context.Context) (*task.Run, error) {
	if c.conn == nil {
		return nil, core.ErrNotConnected
	}
	reply, err := c.tasks.ExecAquire(ctx, &pb.ExecAquireReq{
		Header: c.header(string(c.id)),
	})
	if err != nil {
		return nil, err
	}
	run := pb.UnpackTaskState(reply.Next)
	return &run, nil
}

func (c *executorClient) ExecStop(ctx context.Context) error {
	if c.conn == nil {
		return core.ErrNotConnected
	}
	_, err := c.tasks.ExecStop(ctx, &pb.ExecStopReq{
		Header: c.header(string(c.id)),
	})
	return err
}

func (c *executorClient) Init(ctx context.Context, id task.ID) error {
	if c.conn == nil {
		return core.ErrNotConnected
	}
	_, err := c.tasks.TaskInit(ctx, &pb.TaskInitReq{
		Header:   c.header(string(id)),
		Version:  c.version,
		Executor: string(c.id),
	})
	return err
}

func (c *executorClient) Failure(ctx context.Context, id task.ID, taskErr error) error {
	if c.conn == nil {
		return core.ErrNotConnected
	}
	_, err := c.tasks.TaskFailure(ctx, &pb.TaskFailureReq{
		Header: c.header(string(id)),
		Error:  taskErr.Error(),
	})
	return err
}

func (c *executorClient) Complete(ctx context.Context, id task.ID, result string) error {
	if c.conn == nil {
		return core.ErrNotConnected
	}
	_, err := c.tasks.TaskComplete(ctx, &pb.TaskCompleteReq{
		Header: c.header(string(id)),
		Result: result,
	})
	return err
}

func (c *executorClient) Log(ctx context.Context, id task.ID) (executor.Logger, error) {
	if c.conn == nil {
		return nil, core.ErrNotConnected
	}
	stream, err := c.tasks.TaskLog(ctx)
	if err != nil {
		return nil, err
	}
	return &execLog{
		id:     id,
		client: c,
		stream: stream,
	}, nil
}
