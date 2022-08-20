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

func (s *cowaitServer) QueryClusters(context.Context, *pb.QueryClustersReq) (*pb.QueryClustersReply, error) {
	return &pb.QueryClustersReply{
		Clusters: []*pb.Cluster{
			{
				Name: s.cluster.Name(),
			},
		},
	}, nil
}

func (s *cowaitServer) CreateTask(ctx context.Context, req *pb.CreateTaskReq) (*pb.CreateTaskReply, error) {
	if req.Spec.Cluster != "" && req.Spec.Cluster != s.cluster.Name() {
		return nil, core.ErrUnknownCluster
	}

	def := pb.UnpackTaskSpec(req.Spec)
	instance, err := s.mgr.Schedule(core.TaskID(req.Id), def)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskReply{
		Instance: &pb.Task{
			Id:        string(instance.ID()),
			Spec:      pb.PackTaskSpec(instance.Spec()),
			Status:    string(instance.Status()),
			Scheduled: timestamppb.New(instance.Scheduled()),
			Started:   timestamppb.New(instance.Started()),
			Completed: timestamppb.New(instance.Completed()),
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
