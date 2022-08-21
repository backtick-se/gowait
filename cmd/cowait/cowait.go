package main

import (
	"cowait/adapter/api/grpc"
	"cowait/core"
	"cowait/core/client"
	"fmt"

	"context"

	"go.uber.org/fx"
)

// container client / process manager

func main() {
	executor := fx.New(
		grpc.Module,

		fx.Invoke(run),

		fx.NopLogger,
	)
	executor.Run()
}

func run(ctrl fx.Shutdowner, cli client.API) error {
	hostname := "localhost:50949"
	if err := cli.Connect(hostname); err != nil {
		return err
	}

	ctx := context.Background()
	id, err := cli.CreateTask(ctx, &core.TaskSpec{
		Name:    "client-task",
		Image:   "cowait/gowait-python",
		Command: []string{"python", "-u", "hello.py"},
	})
	if err != nil {
		return err
	}
	fmt.Println("created task", id)
	return ctrl.Shutdown()
}
