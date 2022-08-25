package grpc

import (
	"github.com/backtick-se/gowait/adapter/api/grpc/pb"
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type apiHandler struct {
	pb.UnimplementedCowaitServer

	cluster cluster.T
}

func RegisterApiServer(
	lc fx.Lifecycle,
	srv *grpc.Server,
	cluster cluster.T,
) {
	pb.RegisterCowaitServer(srv, &apiHandler{
		cluster: cluster,
	})
	fmt.Println("registered grpc api server")
}

func (s *apiHandler) CreateTask(ctx context.Context, req *pb.CreateTaskReq) (*pb.CreateTaskReply, error) {
	def := pb.UnpackTaskSpec(req.Spec)
	instance, err := s.cluster.Create(ctx, def)
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

func (s *apiHandler) QueryTasks(ctx context.Context, req *pb.QueryTasksReq) (*pb.QueryTasksReply, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *apiHandler) KillTask(ctx context.Context, req *pb.KillTaskReq) (*pb.KillTaskReply, error) {
	if err := s.cluster.Destroy(ctx, task.ID(req.Id)); err != nil {
		return nil, err
	}
	return &pb.KillTaskReply{}, nil
}

func (s *apiHandler) AwaitTask(req *pb.AwaitTaskReq, server pb.Cowait_AwaitTaskServer) error {
	return fmt.Errorf("not implemented")
}
