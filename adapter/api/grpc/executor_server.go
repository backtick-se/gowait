package grpc

import (
	"cowait/adapter/api/grpc/pb"
	"cowait/core"
	"cowait/core/executor"
	"cowait/core/msg"

	"context"
	"encoding/json"
	"fmt"
	"io"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type executorServer struct {
	pb.UnimplementedExecutorServer
	handler executor.Handler
}

func RegisterExecutorServer(lc fx.Lifecycle, srv *grpc.Server, handler executor.Handler) {
	pb.RegisterExecutorServer(srv, &executorServer{
		handler: handler,
	})
	fmt.Println("registered grpc executor server")
}

func (t *executorServer) TaskInit(ctx context.Context, req *pb.TaskInitReq) (*pb.TaskInitReply, error) {
	err := t.handler.Init(&msg.TaskInit{
		Header:  pb.UnpackHeader(req.Header),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskInitReply{}, nil
}

func (t *executorServer) TaskFailure(ctx context.Context, req *pb.TaskFailureReq) (*pb.TaskFailureReply, error) {
	taskErr := core.NewError(req.Error)
	err := t.handler.Fail(&msg.TaskFailure{
		Header: pb.UnpackHeader(req.Header),
		Error:  taskErr,
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskFailureReply{}, nil
}

func (t *executorServer) TaskComplete(ctx context.Context, req *pb.TaskCompleteReq) (*pb.TaskCompleteReply, error) {
	err := t.handler.Complete(&msg.TaskComplete{
		Header: pb.UnpackHeader(req.Header),
		Result: json.RawMessage(req.Result),
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
		t.handler.Log(&msg.LogEntry{
			Header: pb.UnpackHeader(entry.Header),
			File:   entry.File,
			Data:   entry.Data,
		})
	}
}
