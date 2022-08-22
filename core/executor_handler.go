package core

import (
	"cowait/core/msg"
)

// Handles commands from executors
type ExecutorHandler interface {
	Init(*msg.TaskInit) error
	Complete(*msg.TaskComplete) error
	Fail(*msg.TaskFailure) error
	Log(*msg.LogEntry) error
}
