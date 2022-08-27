package task

import (
	"github.com/backtick-se/gowait/util"
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

type Result []byte

var NoResult = Result("{}")

type Run struct {
	*Spec
	ID        ID
	Parent    ID
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
