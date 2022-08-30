package task

import (
	"time"
)

type Run struct {
	*Spec
	ID        ID
	Parent    ID
	Executor  ID
	Status    Status
	Scheduled time.Time
	Started   time.Time
	Completed time.Time
	Result    Result
	Err       error
}

func NewRun(spec *Spec) *Run {
	return &Run{
		ID:        GenerateID(spec.Image),
		Spec:      spec,
		Status:    StatusWait,
		Scheduled: time.Now(),
	}
}

func (i *Run) init(executor ID) {
	i.Status = StatusExec
	i.Started = time.Now()
	i.Executor = executor
}

func (i *Run) complete(result Result) {
	i.Result = result
	i.Status = StatusDone
	i.Completed = time.Now()
}

func (i *Run) fail(err error) {
	i.Err = err
	i.Status = StatusFail
	i.Completed = time.Now()
}
