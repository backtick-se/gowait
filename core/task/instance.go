package task

import (
	"context"
	"fmt"
	"time"
)

type Instance interface {
	T
	Handler

	Executor() ID

	Exec(done chan struct{})
}

type instance struct {
	run      *Run
	logs     map[string][]string
	executor ID

	on_init     chan *MsgInit
	on_fail     chan *MsgFailure
	on_complete chan *MsgComplete
	on_log      chan *MsgLog
}

func newInstance(spec *Spec) Instance {
	return &instance{
		run: &Run{
			ID:        GenerateID("task"),
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

func (t *instance) ID() ID       { return t.run.ID }
func (t *instance) Spec() *Spec  { return t.run.Spec }
func (t *instance) State() *Run  { return t.run }
func (t *instance) Executor() ID { return t.executor }

func (t *instance) Logs(file string) ([]string, error) {
	return t.logs[file], nil
}

//
// Instance events
//

func (i *instance) OnInit(m *MsgInit) error         { i.on_init <- m; return nil }
func (i *instance) OnFailure(m *MsgFailure) error   { i.on_fail <- m; return nil }
func (i *instance) OnComplete(m *MsgComplete) error { i.on_complete <- m; return nil }
func (i *instance) OnLog(m *MsgLog) error           { i.on_log <- m; return nil }

func (i *instance) Exec(done chan struct{}) {
	ctx := context.Background()
	defer func() {
		fmt.Println("instance", i.run.ID, "exited")
		done <- struct{}{}
		close(done)
	}()
	defer close(i.on_init)
	defer close(i.on_complete)
	defer close(i.on_fail)
	defer close(i.on_log)

	for {
		select {
		case m := <-i.on_init:
			i.executor = m.Executor
			i.run.Init()

		case req := <-i.on_complete:
			i.run.Complete(req.Result)
			return

		case req := <-i.on_fail:
			i.run.Fail(req.Error)
			return

		case req := <-i.on_log:
			log, exists := i.logs[req.File]
			if !exists {
				log = make([]string, 0, 32)
			}
			i.logs[req.File] = append(log, req.Data)

		case <-ctx.Done():
			i.run.Fail(fmt.Errorf("killed by task manager: timeout exceeded"))
			return
		}
	}
}
