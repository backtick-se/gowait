package grpc

import (
	"cowait/core"
	"cowait/core/daemon"
	"cowait/core/msg"
	"cowait/core/pb"

	"context"
	"encoding/json"
	"io"
)

type taskServer struct {
	pb.UnimplementedTaskServer
	mgr daemon.TaskServer
}

func NewTaskServer(mgr daemon.TaskServer) pb.TaskServer {
	return &taskServer{
		mgr: mgr,
	}
}

func (t *taskServer) TaskInit(ctx context.Context, req *pb.TaskInitReq) (*pb.TaskInitReply, error) {
	err := t.mgr.Init(&msg.TaskInit{
		Header:  pb.ParseHeader(req.Header),
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskInitReply{}, nil
}

func (t *taskServer) TaskFailure(ctx context.Context, req *pb.TaskFailureReq) (*pb.TaskFailureReply, error) {
	taskErr := core.NewError(req.Error)
	err := t.mgr.Fail(&msg.TaskFailure{
		Header: pb.ParseHeader(req.Header),
		Error:  taskErr,
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskFailureReply{}, nil
}

func (t *taskServer) TaskComplete(ctx context.Context, req *pb.TaskCompleteReq) (*pb.TaskCompleteReply, error) {
	err := t.mgr.Complete(&msg.TaskComplete{
		Header: pb.ParseHeader(req.Header),
		Result: json.RawMessage(req.Result),
	})
	if err != nil {
		return nil, err
	}
	return &pb.TaskCompleteReply{}, nil
}

func (t *taskServer) TaskLog(stream pb.Task_TaskLogServer) error {
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
		t.mgr.Log(&msg.LogEntry{
			Header: pb.ParseHeader(entry.Header),
			File:   entry.File,
			Data:   entry.Data,
		})
	}
}
