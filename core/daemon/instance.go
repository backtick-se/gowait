package daemon

import (
	"context"
	"fmt"
	"time"

	"github.com/backtick-se/gowait/core/msg"
	"github.com/backtick-se/gowait/core/task"
)

type instance struct {
	run  *task.Run
	logs map[string][]string

	on_init     chan *msg.TaskInit
	on_fail     chan *msg.TaskFailure
	on_complete chan *msg.TaskComplete
	on_log      chan *msg.LogEntry
}

var _ task.T = &instance{}

func newInstance(spec *task.Spec) *instance {
	return &instance{
		run: &task.Run{
			ID:        task.GenerateID("task"),
			Spec:      spec,
			Status:    task.StatusWait,
			Scheduled: time.Now(),
		},
		logs:        make(map[string][]string),
		on_init:     make(chan *msg.TaskInit),
		on_fail:     make(chan *msg.TaskFailure),
		on_complete: make(chan *msg.TaskComplete),
		on_log:      make(chan *msg.LogEntry),
	}
}

func (t *instance) ID() task.ID {
	return t.run.ID
}

func (t *instance) Spec() *task.Spec {
	return t.run.Spec
}

func (t *instance) State() *task.Run {
	return t.run
}

func (t *instance) Logs(file string) ([]string, error) {
	return t.logs[file], nil
}

func (i *instance) exec(done chan struct{}) {
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
		case <-i.on_init:
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
