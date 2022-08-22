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
	fmt.Println("cluster.Info() call!")
	return &pb.ClusterInfoReply{
		Name: s.cluster.Name(),
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

func (s *clusterHandler) Subscribe(*pb.ClusterSubscribeReq, pb.Cluster_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
