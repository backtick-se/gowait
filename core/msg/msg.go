package msg

import (
	"cowait/core"
	"encoding/json"
	"time"
)

type Header struct {
	ID   core.TaskID
	Time time.Time
}

type TaskInit struct {
	Header
	Version string
}

type TaskFailure struct {
	Header
	Error error
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
