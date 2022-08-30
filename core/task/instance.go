package task

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var ErrInactiveTask = errors.New("task is inactive")

type Instance interface {
	T
	Handler

	Exec(done chan struct{})
}

type instance struct {
	run    *Run
	logs   map[string][]string
	active bool

	on_init     chan *MsgInit
	on_fail     chan *MsgFailure
	on_complete chan *MsgComplete
	on_log      chan *MsgLog
}

func NewInstance(spec *Spec) Instance {
	return &instance{
		run: &Run{
			ID:        GenerateID(spec.Image),
			Spec:      spec,
			Status:    StatusWait,
			Scheduled: time.Now(),
		},
		logs:        make(map[string][]string),
		on_init:     make(chan *MsgInit),
		on_fail:     make(chan *MsgFailure),
		on_complete: make(chan *MsgComplete),
		on_log:      make(chan *MsgLog),
	}
}

func (t *instance) ID() ID      { return t.run.ID }
func (t *instance) Spec() *Spec { return t.run.Spec }
func (t *instance) State() *Run { return t.run } // maybe return a copy instead

func (t *instance) Logs(file string) ([]string, error) {
	return t.logs[file], nil
}

func (i *instance) Exec(done chan struct{}) {
	if i.active {
		panic("task routine is already running")
	}
	i.active = true
	defer close(i.on_init)
	defer close(i.on_complete)
	defer close(i.on_fail)
	defer close(i.on_log)
	defer func() {
		i.active = false
		fmt.Println("instance", i.run.ID, "exited")
		done <- struct{}{}
		close(done)
	}()

	ctx := context.Background()
	// todo: timeout

	for {
		select {
		case m := <-i.on_init:
			i.run.init(m.Executor)

		case req := <-i.on_complete:
			i.run.complete(req.Result)
			return

		case req := <-i.on_fail:
			i.run.fail(req.Error)
			return

		case req := <-i.on_log:
			log, exists := i.logs[req.File]
			if !exists {
				log = make([]string, 0, 32)
			}
			i.logs[req.File] = append(log, req.Data)

		case <-ctx.Done():
			i.run.fail(fmt.Errorf("killed by task manager: timeout exceeded"))
			return
		}
	}
}

//
// Instance events
//

func (i *instance) OnInit(m *MsgInit) error {
	if !i.active {
		return ErrInactiveTask
	}
	i.on_init <- m
	return nil
}

func (i *instance) OnFailure(m *MsgFailure) error {
	if !i.active {
		return ErrInactiveTask
	}
	i.on_fail <- m
	return nil
}

func (i *instance) OnComplete(m *MsgComplete) error {
	if !i.active {
		return ErrInactiveTask
	}
	i.on_complete <- m
	return nil
}

func (i *instance) OnLog(m *MsgLog) error {
	if !i.active {
		return ErrInactiveTask
	}
	i.on_log <- m
	return nil
}
