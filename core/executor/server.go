package executor

import (
	"context"
	"cowait/adapter/api/grpc"
	"cowait/adapter/api/grpc/pb"
	"cowait/core"
	"cowait/core/msg"

	"fmt"
	"net"

	"go.uber.org/fx"
	grpcio "google.golang.org/grpc"
)

type Server interface {
	Listen(port int) error
	Close() error

	OnInit() <-chan *msg.TaskInit
	OnComplete() <-chan *msg.TaskComplete
	OnFailure() <-chan *msg.TaskFailure
	OnLog() <-chan *msg.LogEntry
}

type server struct {
	grpc *grpcio.Server

	on_init     chan *msg.TaskInit
	on_complete chan *msg.TaskComplete
	on_failure  chan *msg.TaskFailure
	on_log      chan *msg.LogEntry
}

var _ core.ExecutorHandler = &server{}

func NewServer(lc fx.Lifecycle) Server {
	// its not really OK to import grpc stuff here.
	// todo: extract

	server := &server{
		grpc: grpcio.NewServer(),

		on_init:     make(chan *msg.TaskInit),
		on_complete: make(chan *msg.TaskComplete),
		on_failure:  make(chan *msg.TaskFailure),
		on_log:      make(chan *msg.LogEntry),
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

func (s *server) OnInit() <-chan *msg.TaskInit         { return s.on_init }
func (s *server) OnComplete() <-chan *msg.TaskComplete { return s.on_complete }
func (s *server) OnFailure() <-chan *msg.TaskFailure   { return s.on_failure }
func (s *server) OnLog() <-chan *msg.LogEntry          { return s.on_log }

func (s *server) Close() error {
	defer close(s.on_init)
	defer close(s.on_complete)
	defer close(s.on_failure)
	defer close(s.on_log)

	s.grpc.GracefulStop()

	return nil
}

func (t *server) Init(req *msg.TaskInit) error {
	t.on_init <- req
	return nil
}

func (t *server) Complete(req *msg.TaskComplete) error {
	t.on_complete <- req
	return nil
}

func (t *server) Fail(req *msg.TaskFailure) error {
	t.on_failure <- req
	return nil
}

func (t *server) Log(req *msg.LogEntry) error {
	t.on_log <- req
	return nil
}
