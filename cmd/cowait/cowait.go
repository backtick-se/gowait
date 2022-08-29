package main

import (
	"github.com/backtick-se/gowait/adapter/api/grpc"
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"

	"go.uber.org/fx"
)

// container client / process manager

func main() {
	executor := fx.New(
		grpc.Module,

		fx.Invoke(run),

		// fx.NopLogger,
	)
	executor.Run()
}

func run(shut fx.Shutdowner, cli core.APIClient) error {
	hostname := "localhost:1337"
	if err := cli.Connect(hostname); err != nil {
		return err
	}

	ctx := context.Background()
	id, err := cli.CreateTask(ctx, &task.Spec{
		Name:    "cowait.builtin.enumerate",
		Image:   "cowait/gowait-python",
		Command: []string{"python", "-um", "cowait"},
	})
	if err != nil {
		return err
	}
	fmt.Println("created task", id)
	return shut.Shutdown()
}
