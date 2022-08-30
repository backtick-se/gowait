package task

import (
	"context"
	"time"
)

type Instance interface {
	T
	Handler

	// Starts the instance handler goroutine
	Start() chan struct{}

	// Explicitly triggers an OnFailure event
	Fail(err error) error
}

type instance struct {
	run    *Run
	logs   map[string][]string
	active bool

	on_init     chan *MsgInit
	on_fail     chan *MsgFailure
	on_complete chan *MsgComplete
	on_log      chan *MsgLog
	err         chan error
}

func NewInstance(spec *Spec) Instance {
	return &instance{
		run:         NewRun(spec),
		logs:        make(map[string][]string),
		on_init:     make(chan *MsgInit),
		on_fail:     make(chan *MsgFailure),
		on_complete: make(chan *MsgComplete),
		on_log:      make(chan *MsgLog),
		err:         make(chan error),
	}
}

func (t *instance) ID() ID      { return t.run.ID }
func (t *instance) Spec() *Spec { return t.run.Spec }
func (t *instance) State() *Run { return t.run } // maybe return a copy instead

func (t *instance) Logs(file string) ([]string, error) {
	return t.logs[file], nil
}

func (i *instance) Start() chan struct{} {
	if i.active {
		panic("task routine is already running")
	}
	done := make(chan struct{})
	go i.proc(done)
	<-i.err
	return done
}

func (i *instance) Fail(err error) error {
	return i.OnFailure(&MsgFailure{
		Header: Header{
			ID:   string(i.run.ID),
			Time: time.Now(),
		},
		Error: err,
	})
}

func (i *instance) proc(done chan struct{}) {
	i.active = true
	i.err <- nil

	defer close(i.on_init)
	defer close(i.on_complete)
	defer close(i.on_fail)
	defer close(i.on_log)
	defer func() {
		i.active = false
		done <- struct{}{}
		close(done)
	}()

	ctx := context.Background()
	// todo: timeout

	for {
		select {
		case m := <-i.on_init:
			i.run.init(m.Executor)
			i.err <- nil

		case req := <-i.on_complete:
			i.run.complete(req.Result)
			i.err <- nil
			return

		case req := <-i.on_fail:
			i.run.fail(req.Error)
			i.err <- nil
			return

		case req := <-i.on_log:
			log, exists := i.logs[req.File]
			if !exists {
				log = make([]string, 0, 32)
			}
			i.logs[req.File] = append(log, req.Data)
			i.err <- nil

		case <-ctx.Done():
			i.run.fail(ErrTimeout)
			return
		}
	}
}

//
// Instance events
//

func (i *instance) OnInit(m *MsgInit) error {
	if !i.active {
		return ErrInactive
	}
	i.on_init <- m
	return <-i.err
}

func (i *instance) OnFailure(m *MsgFailure) error {
	if !i.active {
		return ErrInactive
	}
	i.on_fail <- m
	return <-i.err
}

func (i *instance) OnComplete(m *MsgComplete) error {
	if !i.active {
		return ErrInactive
	}
	i.on_complete <- m
	return <-i.err
}

func (i *instance) OnLog(m *MsgLog) error {
	if !i.active {
		return ErrInactive
	}
	i.on_log <- m
	return <-i.err
}
