package k8s

import "cowait/core"

type task struct {
	job string
}

var _ core.Task = &task{}
