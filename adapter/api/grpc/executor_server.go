package grpc

import (
	"github.com/backtick-se/gowait/adapter/api/grpc/pb"
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"
	"github.com/backtick-se/gowait/util/slices"

	"context"
	"fmt"
	"io"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type executorServer struct {
	pb.UnimplementedExecutorServer
	handler executor.Handler
	tasks   task.Handler
}

func RegisterExecutorServer(lc fx.Lifecycle, srv *grpc.Server, handler executor.Handler, tasks task.Handler) {
	pb.RegisterExecutorServer(srv, &executorServer{
		handler: handler,
		tasks:   tasks,
	})
	fmt.Println("registered grpc executor server")
}

func (t *executorServer) ExecInit(ctx context.Context, req *pb.ExecInitReq) (*pb.ExecInitReply, error) {
	err := t.handler.ExecInit(ctx, &executor.MsgInit{
		Header: pb.UnpackExecutorHeader(req.Header),
		Image:  req.Image,
		Specs:  slices.Map(req.Specs, pb.UnpackTaskSpec),
	})
	return &pb.ExecInitReply{}, err
}

func (t *executorServer) ExecAquire(ctx context.Context, req *pb.ExecAquireReq) (*pb.ExecAquireReply, error) {
	spec, err := t.handler.ExecAquire(ctx, &executor.MsgAquire{
		Header: pb.UnpackExecutorHeader(req.Header),
	})
	if err != nil {
		return nil, err
	}
	if spec == nil {
		return &pb.ExecAquireReply{
			Next: nil,
		}, nil
	}

	return &pb.ExecAquireReply{
		Next: pb.PackTaskState(spec),
	}, nil
}

func (t *executorServer) ExecStop(ctx context.Context, req *pb.ExecStopReq) (*pb.ExecStopReply, error) {
	err := t.handler.ExecStop(ctx, &executor.MsgStop{
		Header: pb.UnpackExecutorHeader(req.Header),
	})
	return &pb.ExecStopReply{}, err
}

//
// Task API
//

func (t *executorServer) TaskInit(ctx context.Context, req *pb.TaskInitReq) (*pb.TaskInitReply, error) {
	err := t.tasks.OnInit(&task.MsgInit{
		Header:   pb.UnpackTaskHeader(req.Header),
		Version:  req.Version,
		Executor: task.ID(req.Executor),
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskInitReply{}, nil
}

func (t *executorServer) TaskFailure(ctx context.Context, req *pb.TaskFailureReq) (*pb.TaskFailureReply, error) {
	taskErr := core.NewError(req.Error)
	err := t.tasks.OnFailure(&task.MsgFailure{
		Header: pb.UnpackTaskHeader(req.Header),
		Error:  taskErr,
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskFailureReply{}, nil
}

func (t *executorServer) TaskComplete(ctx context.Context, req *pb.TaskCompleteReq) (*pb.TaskCompleteReply, error) {
	err := t.tasks.OnComplete(&task.MsgComplete{
		Header: pb.UnpackTaskHeader(req.Header),
		Result: task.Result(req.Result),
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskCompleteReply{}, nil
}

func (t *executorServer) TaskLog(stream pb.Executor_TaskLogServer) error {
	records := 0
	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.LogSummary{
				Records: int64(records),
			})
		}
		if err != nil {
			return err
		}
		records++
		t.tasks.OnLog(&task.MsgLog{
			Header: pb.UnpackTaskHeader(entry.Header),
			File:   entry.File,
			Data:   entry.Data,
		})
	}
}
