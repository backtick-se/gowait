package core

import "errors"

var ErrUnknownTask = errors.New("unknown task")

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
