package grpc

import (
	"cowait/adapter/api/grpc/pb"

	"context"
	"fmt"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewServer(
	lc fx.Lifecycle,
	taskServer pb.ExecutorServer,
	cowaitServer pb.CowaitServer,
) {
	port := 1337

	grpcServer := grpc.NewServer()
	pb.RegisterExecutorServer(grpcServer, taskServer)
	pb.RegisterCowaitServer(grpcServer, cowaitServer)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
			if err != nil {
				return fmt.Errorf("failed to listen: %v", err)
			}

			go grpcServer.Serve(listener)
			return nil
		},
	})
}
