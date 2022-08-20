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
	if req.Task.Cluster != "" && req.Task.Cluster != s.cluster.Name() {
		return nil, core.ErrUnknownCluster
	}

	def := pb.UnpackTaskdef(req.Task)
	instance, err := s.mgr.Schedule(def)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskReply{
		Instance: &pb.Instance{
			Id:        string(instance.ID()),
			Task:      req.Task,
			Status:    string(instance.Status()),
			Scheduled: timestamppb.New(instance.Scheduled()),
		},
	}, nil
}

func (s *cowaitServer) QueryTasks(context.Context, *pb.QueryTasksReq) (*pb.QueryTasksReply, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *cowaitServer) KillTask(context.Context, *pb.KillTaskReq) (*pb.KillTaskReply, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *cowaitServer) AwaitTask(req *pb.AwaitTaskReq, server pb.Cowait_AwaitTaskServer) error {
	return fmt.Errorf("not implemented")
}
