package msg

import (
	"time"

	"github.com/backtick-se/gowait/core/task"
)

type Header struct {
	ID   string
	Time time.Time
}

type ExecInit struct {
	Header
	Version string
	Image   string
	Specs   []*task.Spec
}

type TaskInit struct {
	Header
	Version  string
	Executor task.ID
}

type ExecAquire struct {
	Header
}

type ExecStop struct {
	Header
}

type TaskFailure struct {
	Header
	Error error
}

type TaskComplete struct {
	Header
	Result task.Result
}

type LogEntry struct {
	Header
	File string
	Data string
}
