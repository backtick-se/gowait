package grpc

import (
	"cowait/adapter/api/grpc/pb"
	"cowait/core"

	"fmt"
	"net"

	"github.com/hashicorp/yamux"
	"google.golang.org/grpc"
)

type uplink struct {
	cluster core.Cluster
	srv     *grpc.Server
}

func NewUplinkClient(cluster core.Cluster) core.UplinkClient {
	return &uplink{
		cluster: cluster,
	}
}

func (u *uplink) Connect(endpoint string) error {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		fmt.Println("failed to dial", endpoint, ":", err)
		return err
	}

	listener, err := yamux.Server(conn, yamux.DefaultConfig())
	if err != nil {
		fmt.Println("failed to create yamux listener:", err)
		return err
	}

	// create a gRPC server object
	u.srv = grpc.NewServer()
	pb.RegisterClusterServer(u.srv, NewClusterHandler(u.cluster))

	if err := u.srv.Serve(listener); err != nil {
		fmt.Println("uplink client failed:", err)
		return err
	}

	return nil
}

func (u *uplink) Close() error {
	if u.srv != nil {
		u.srv.GracefulStop()
		u.srv = nil
	}
	return nil
}
