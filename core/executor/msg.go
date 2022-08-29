package executor

import (
	"time"

	"github.com/backtick-se/gowait/core/task"
)

type Header struct {
	ID   string
	Time time.Time
}

type MsgInit struct {
	Header
	Version string
	Image   string
	Specs   []*task.Spec
}
type MsgAquire struct {
	Header
}

type MsgStop struct {
	Header
}
