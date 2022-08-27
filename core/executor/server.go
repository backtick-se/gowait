package executor

import (
	"github.com/backtick-se/gowait/core/msg"
	"github.com/backtick-se/gowait/core/task"

	"go.uber.org/fx"
)

type Server interface {
	Handler

	Close() error
	OnInit() <-chan *msg.TaskInit
	OnComplete() <-chan *msg.TaskComplete
	OnFailure() <-chan *msg.TaskFailure
	OnLog() <-chan *msg.LogEntry
}

type server struct {
	on_init     chan *msg.TaskInit
	on_complete chan *msg.TaskComplete
	on_failure  chan *msg.TaskFailure
	on_log      chan *msg.LogEntry
}

func NewServer(lc fx.Lifecycle) Server {
	server := &server{
		on_init:     make(chan *msg.TaskInit),
		on_complete: make(chan *msg.TaskComplete),
		on_failure:  make(chan *msg.TaskFailure),
		on_log:      make(chan *msg.LogEntry),
	}

	return server
}

func registerExecutorHandler(server Server) Handler {
	return server
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
	return nil
}

func (t *server) Init(req *msg.TaskInit) error {
	t.on_init <- req
	return nil
}

func (t *server) ExecInit(*msg.ExecInit) error {
	return nil
}

func (t *server) ExecAquire(*msg.ExecAquire) (*task.Run, error) {
	return nil, nil
}

func (t *server) ExecStop(*msg.ExecStop) error {
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
