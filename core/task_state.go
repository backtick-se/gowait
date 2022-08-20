package core

import (
	"cowait/util"
	"encoding/json"
	"time"
)

type TaskID string
type TaskStatus string

const None = TaskID("")

const (
	StatusWait TaskStatus = "wait"
	StatusExec TaskStatus = "exec"
	StatusFail TaskStatus = "fail"
	StatusDone TaskStatus = "done"
)

func GenerateTaskID(name string) TaskID {
	return TaskID(name + "-" + util.RandomString(6))
}

type TaskState struct {
	ID        TaskID
	Parent    TaskID
	Status    TaskStatus
	Spec      *TaskSpec
	Scheduled time.Time
	Started   time.Time
	Completed time.Time
	Result    json.RawMessage
	Err       error
}

func (i *TaskState) Init() {
	i.Status = StatusExec
	i.Started = time.Now()
}

func (i *TaskState) Complete(result json.RawMessage) {
	i.Result = result
	i.Status = StatusDone
	i.Completed = time.Now()
}

func (i *TaskState) Fail(err error) {
	i.Err = err
	i.Status = StatusFail
	i.Completed = time.Now()
}
