package msg

import (
	"cowait/core"
	"encoding/json"
	"time"
)

type Header struct {
	ID      core.TaskID
	Version string
	Time    time.Time
}

type TaskInit struct {
	Header
	Task core.TaskDef
}

type TaskStatus struct {
	Header
	Status core.TaskStatus
}

type TaskFailure struct {
	Header
	Error string
}

type TaskComplete struct {
	Header
	Result json.RawMessage
}

type LogEntry struct {
	Header
	File string
	Data string
}
