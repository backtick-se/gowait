package task

import (
	"time"
)

type Header struct {
	ID   string
	Time time.Time
}

type MsgInit struct {
	Header
	Version  string
	Executor ID
}

type MsgFailure struct {
	Header
	Error error
}

type MsgComplete struct {
	Header
	Result Result
}

type MsgLog struct {
	Header
	File string
	Data string
}
