package core

type TaskID string
type TaskStatus string

const (
	StatusWait TaskStatus = "wait"
	StatusExec TaskStatus = "exec"
	StatusFail TaskStatus = "fail"
	StatusDone TaskStatus = "done"
)

type Task interface {
}
