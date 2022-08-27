package task

import (
	"time"
)

type MsgHeader struct {
	ID   string
	Time time.Time
}

type MsgInit struct {
	MsgHeader
	Version  string
	Executor ID
}

type MsgFailure struct {
	MsgHeader
	Error error
}

type MsgComplete struct {
	MsgHeader
	Result Result
}

type MsgLog struct {
	MsgHeader
	File string
	Data string
}
