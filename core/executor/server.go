package executor

import (
	"context"

	"github.com/backtick-se/gowait/core/task"

	"go.uber.org/fx"
)

type Server interface {
	Handler
	task.Handler

	Close() error
	SetRun(*task.Run)

	AwaitInit() <-chan *task.MsgInit
	AwaitComplete() <-chan *task.MsgComplete
	AwaitFailure() <-chan *task.MsgFailure
	AwaitLog() <-chan *task.MsgLog
}

type server struct {
	run         *task.Run
	on_init     chan *task.MsgInit
	on_complete chan *task.MsgComplete
	on_failure  chan *task.MsgFailure
	on_log      chan *task.MsgLog
}

func NewServer(lc fx.Lifecycle) Server {
	server := &server{
		on_init:     make(chan *task.MsgInit),
		on_complete: make(chan *task.MsgComplete),
		on_failure:  make(chan *task.MsgFailure),
		on_log:      make(chan *task.MsgLog),
	}
	return server
}

func (s *server) AwaitInit() <-chan *task.MsgInit         { return s.on_init }
func (s *server) AwaitComplete() <-chan *task.MsgComplete { return s.on_complete }
func (s *server) AwaitFailure() <-chan *task.MsgFailure   { return s.on_failure }
func (s *server) AwaitLog() <-chan *task.MsgLog           { return s.on_log }

func (s *server) SetRun(run *task.Run) {
	s.run = run
}

func (s *server) Close() error {
	defer close(s.on_init)
	defer close(s.on_complete)
	defer close(s.on_failure)
	defer close(s.on_log)
	return nil
}

func (t *server) ExecInit(context.Context, *MsgInit) error {
	return nil
}

func (t *server) ExecAquire(context.Context, *MsgAquire) (*task.Run, error) {
	return t.run, nil
}

func (t *server) ExecStop(context.Context, *MsgStop) error {
	return nil
}

func (t *server) OnInit(req *task.MsgInit) error {
	t.on_init <- req
	return nil
}

func (t *server) OnComplete(req *task.MsgComplete) error {
	t.on_complete <- req
	return nil
}

func (t *server) OnFailure(req *task.MsgFailure) error {
	t.on_failure <- req
	return nil
}

func (t *server) OnLog(req *task.MsgLog) error {
	t.on_log <- req
	return nil
}
