package executor

type Status string

const (
	StatusWait  Status = "wait"
	StatusIdle  Status = "idle"
	StatusExec  Status = "exec"
	StatusStop  Status = "stop"
	StatusCrash Status = "crash"
)
