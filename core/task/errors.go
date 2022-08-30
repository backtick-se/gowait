package task

import "errors"

var ErrTimeout = errors.New("timeout")
var ErrCanceled = errors.New("canceled manually")
