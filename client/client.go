package client

import (
	"cowait/core"

	"fmt"
)

type Executor interface {
	Run() error
}

type executor struct {
	task *core.TaskDef
}

func ExecutorFromEnv(envdef string) (Executor, error) {
	def, err := core.TaskDefFromEnv(envdef)
	if err != nil {
		return nil, err
	}

	return &executor{
		task: def,
	}, nil
}

func (e *executor) Run() error {
	fmt.Printf("running task: %+v\n", e.task)
	return nil
}
