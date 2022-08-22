package daemon

import (
	"cowait/adapter/api/grpc"
	"cowait/adapter/api/grpc/pb"
	"cowait/core"
	"fmt"

	"context"
	"net"

	"github.com/hashicorp/yamux"
	"go.uber.org/fx"
	grpcio "google.golang.org/grpc"
)

type Uplink interface {
	Connect(endpoint string) error
	Push()
}

type uplink struct {
	cluster core.Cluster
}

func NewUplink(lc fx.Lifecycle, cluster core.Cluster) Uplink {
	u := &uplink{
		cluster: cluster,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go u.Connect("cloud.default.svc.cluster.local:1338")
			return nil
		},
	})
	return u
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
	grpcServer := grpcio.NewServer()
	pb.RegisterClusterServer(grpcServer, grpc.NewClusterHandler(u.cluster))

	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("uplink client failed:", err)
		return err
	}

	return nil
}

func (u *uplink) Push() {}
