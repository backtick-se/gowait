package grpc

import (
	"cowait/core"
	"cowait/core/daemon"
	"cowait/core/pb"

	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type cowaitServer struct {
	pb.UnimplementedCowaitServer

	cluster core.Cluster
	mgr     daemon.TaskManager
}

func NewCowaitServer(cluster core.Cluster, mgr daemon.TaskManager) pb.CowaitServer {
	return &cowaitServer{
		cluster: cluster,
		mgr:     mgr,
	}
}

func (s *cowaitServer) CreateTask(ctx context.Context, req *pb.CreateTaskReq) (*pb.CreateTaskReply, error) {
	def := pb.UnpackTaskSpec(req.Spec)
	instance, err := s.mgr.Schedule(def)
	if err != nil {
		return nil, err
	}

	state := instance.State()
	return &pb.CreateTaskReply{
		Task: &pb.Task{
			TaskId:    string(instance.ID()),
			Spec:      pb.PackTaskSpec(instance.Spec()),
			Status:    string(state.Status),
			Scheduled: timestamppb.New(state.Scheduled),
			Started:   timestamppb.New(state.Started),
			Completed: timestamppb.New(state.Completed),
		},
	}, nil
}

func (s *cowaitServer) QueryTasks(ctx context.Context, req *pb.QueryTasksReq) (*pb.QueryTasksReply, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *cowaitServer) KillTask(ctx context.Context, req *pb.KillTaskReq) (*pb.KillTaskReply, error) {
	if err := s.cluster.Kill(ctx, core.TaskID(req.Id)); err != nil {
		return nil, err
	}
	return &pb.KillTaskReply{}, nil
}

func (s *cowaitServer) AwaitTask(req *pb.AwaitTaskReq, server pb.Cowait_AwaitTaskServer) error {
	return fmt.Errorf("not implemented")
}
