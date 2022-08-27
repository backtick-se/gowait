package task

import "time"

type T interface {
	ID() ID
	Spec() *Spec
	State() *Run
	Logs(file string) ([]string, error)
}

type task struct {
	run  *Run
	logs map[string][]string
}

func New(spec *Spec) T {
	return &task{
		run: &Run{
			ID:        GenerateID("task"),
			Spec:      spec,
			Status:    StatusWait,
			Scheduled: time.Now(),
		},
		logs: make(map[string][]string),
	}
}

func (t *task) ID() ID {
	return t.run.ID
}

func (t *task) Spec() *Spec {
	return t.run.Spec
}

func (t *task) State() *Run {
	return t.run
}

func (t *task) Logs(file string) ([]string, error) {
	return t.logs[file], nil
}
