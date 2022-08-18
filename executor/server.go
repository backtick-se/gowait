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
}

var _ daemon.TaskServer = &server{}

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

func newServer() (*server, error) {
	port := 1337
	server := &server{
		completed: make(chan *msg.TaskComplete),
		failed:    make(chan *msg.TaskFailure),
		log:       make(chan *msg.LogEntry),
	}

	grpcServer := grpcio.NewServer()
	pb.RegisterTaskServer(grpcServer, grpc.NewTaskServer(server))

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	go grpcServer.Serve(listener)

	return server, nil
}
