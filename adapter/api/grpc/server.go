package grpc

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewServer(lc fx.Lifecycle) *grpc.Server {
	port := 1337
	srv := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
			if err != nil {
				return fmt.Errorf("failed to listen: %v", err)
			}
			fmt.Println("grpc listening on port", port)

			go srv.Serve(listener)
			return nil
		},

		OnStop: func(context.Context) error {
			srv.GracefulStop()
			return nil
		},
	})

	return srv
}
