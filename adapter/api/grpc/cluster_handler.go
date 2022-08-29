package grpc

import (
	"github.com/backtick-se/gowait/adapter/api/grpc/pb"
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type clusterHandler struct {
	pb.UnimplementedClusterServer
	cluster cluster.T
}

func NewClusterHandler(cluster cluster.T) pb.ClusterServer {
	return &clusterHandler{
		cluster: cluster,
	}
}

func (s *clusterHandler) Info(context.Context, *pb.ClusterInfoReq) (*pb.ClusterInfoReply, error) {
	info := s.cluster.Info()
	return &pb.ClusterInfoReply{
		Name: info.Name,
	}, nil
}

func (s *clusterHandler) Spawn(context.Context, *pb.CreateTaskReq) (*pb.CreateTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Spawn not implemented")
}

func (s *clusterHandler) Kill(ctx context.Context, req *pb.KillTaskReq) (*pb.KillTaskReply, error) {
	err := s.cluster.Destroy(ctx, task.ID(req.Id))
	return &pb.KillTaskReply{}, err
}

func (s *clusterHandler) Poke(context.Context, *pb.ClusterPokeReq) (*pb.ClusterPokeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Poke not implemented")
}

func (s *clusterHandler) Subscribe(req *pb.ClusterSubscribeReq, stream pb.Cluster_SubscribeServer) error {
	info := s.cluster.Info()
	sub := s.cluster.Events().Subscribe()
	for {
		event, ok := <-sub.Next()
		if !ok {
			fmt.Println("end of event stream")
			break
		}
		if err := stream.Send(&pb.ClusterEvent{
			ClusterId: info.ID,
			Type:      event.Type,
			Task:      pb.PackTaskState(&event.Task),
		}); err != nil {
			return err
		}
	}
	return nil
}
