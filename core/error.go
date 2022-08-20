package core

import "errors"

var ErrUnknownTask = errors.New("unknown task")
var ErrUnknownCluster = errors.New("unknown cluster")

type TaskError struct {
	message string
}

func NewError(msg string) error {
	return &TaskError{
		message: msg,
	}
}

func (e *TaskError) Error() string {
	return e.message
}
