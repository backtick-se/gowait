package executor

import (
	"context"
	"cowait/adapter/api/grpc"
	"cowait/adapter/api/grpc/pb"
	"cowait/core/daemon"
	"cowait/core/msg"

	"fmt"
	"net"

	"go.uber.org/fx"
	grpcio "google.golang.org/grpc"
)

type Server interface {
	Listen(port int) error
	Close() error
	Completed() <-chan *msg.TaskComplete
	Failed() <-chan *msg.TaskFailure
	Logs() <-chan *msg.LogEntry
}

type server struct {
	grpc *grpcio.Server

	completed chan *msg.TaskComplete
	failed    chan *msg.TaskFailure
	log       chan *msg.LogEntry
}

var _ daemon.ExecutorServer = &server{}

func NewServer(lc fx.Lifecycle) Server {
	// its not really OK to import grpc stuff here.
	// todo: extract

	server := &server{
		grpc: grpcio.NewServer(),

		completed: make(chan *msg.TaskComplete),
		failed:    make(chan *msg.TaskFailure),
		log:       make(chan *msg.LogEntry),
	}

	pb.RegisterExecutorServer(server.grpc, grpc.NewExecutorServer(server))

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return server.Close()
		},
	})

	return server
}

func (s *server) Listen(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on %d: %v", port, err)
	}

	go s.grpc.Serve(listener)
	return nil
}

func (s *server) Completed() <-chan *msg.TaskComplete { return s.completed }
func (s *server) Failed() <-chan *msg.TaskFailure     { return s.failed }
func (s *server) Logs() <-chan *msg.LogEntry          { return s.log }

func (s *server) Close() error {
	defer close(s.completed)
	defer close(s.failed)
	defer close(s.log)

	s.grpc.GracefulStop()

	return nil
}

func (t *server) Init(req *msg.TaskInit) error {
	return nil
}

func (t *server) Complete(req *msg.TaskComplete) error {
	t.completed <- req
	return nil
}

func (t *server) Fail(req *msg.TaskFailure) error {
	t.failed <- req
	return nil
}

func (t *server) Log(req *msg.LogEntry) error {
	t.log <- req
	return nil
}
