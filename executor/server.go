package executor

import (
	"cowait/adapter/api/grpc"
	"cowait/core/daemon"
	"cowait/core/msg"
	"cowait/core/pb"

	"fmt"
	"net"

	grpcio "google.golang.org/grpc"
)

type server struct {
	completed chan *msg.TaskComplete
	failed    chan *msg.TaskFailure
	log       chan *msg.LogEntry

	grpc *grpcio.Server
}

var _ daemon.TaskServer = &server{}

func newServer() (*server, error) {
	port := 1337
	server := &server{
		grpc: grpcio.NewServer(),

		completed: make(chan *msg.TaskComplete),
		failed:    make(chan *msg.TaskFailure),
		log:       make(chan *msg.LogEntry),
	}

	pb.RegisterTaskServer(server.grpc, grpc.NewTaskServer(server))

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	go server.grpc.Serve(listener)

	return server, nil
}

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
