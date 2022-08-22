package grpc

import (
	"cowait/adapter/api/grpc/pb"
	"cowait/core"
	"fmt"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type clusterHandler struct {
	pb.UnimplementedClusterServer
	cluster core.Cluster
}

func NewClusterHandler(cluster core.Cluster) pb.ClusterServer {
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

func (s *clusterHandler) Spawn(context.Context, *pb.ClusterSpawnReq) (*pb.ClusterSpawnReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Spawn not implemented")
}

func (s *clusterHandler) Kill(context.Context, *pb.ClusterKillReq) (*pb.ClusterKillReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kill not implemented")
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
