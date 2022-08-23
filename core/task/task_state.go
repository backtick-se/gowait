package task

import (
	"cowait/util"
	"encoding/json"
	"time"
)

type ID string
type Status string

const None = ID("")

const (
	StatusWait Status = "wait"
	StatusExec Status = "exec"
	StatusFail Status = "fail"
	StatusDone Status = "done"
)

func GenerateID(name string) ID {
	return ID(name + "-" + util.RandomString(6))
}

type State struct {
	ID        ID
	Parent    ID
	Status    Status
	Spec      *Spec
	Scheduled time.Time
	Started   time.Time
	Completed time.Time
	Result    json.RawMessage
	Err       error
}

func (i *State) Init() {
	i.Status = StatusExec
	i.Started = time.Now()
}

func (i *State) Complete(result json.RawMessage) {
	i.Result = result
	i.Status = StatusDone
	i.Completed = time.Now()
}

func (i *State) Fail(err error) {
	i.Err = err
	i.Status = StatusFail
	i.Completed = time.Now()
}
