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

func (i *Run) Init() {
	i.Status = StatusExec
	i.Started = time.Now()
}

func (i *Run) Complete(result Result) {
	i.Result = result
	i.Status = StatusDone
	i.Completed = time.Now()
}

func (i *Run) Fail(err error) {
	i.Err = err
	i.Status = StatusFail
	i.Completed = time.Now()
}
