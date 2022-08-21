package daemon

import (
	"cowait/core/msg"
)

// Handles commands from executors
type ExecutorServer interface {
	Init(*msg.TaskInit) error
	Complete(*msg.TaskComplete) error
	Fail(*msg.TaskFailure) error
	Log(*msg.LogEntry) error
}
