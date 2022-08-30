package task

import (
	"github.com/backtick-se/gowait/util"

	"github.com/gosimple/slug"
)

type ID string
type Status string

const None = ID("")

const (
	StatusWait Status = "wait"
	StatusExec Status = "exec"
	StatusFail Status = "fail"
	StatusDone Status = "done"
)

func GenerateID(name string) ID {
	slugified := slug.Make(name)
	return ID(slugified + "-" + util.RandomString(6))
}

type Result []byte

var NoResult = Result("{}")
