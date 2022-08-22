package main

import (
	"cowait/adapter/api/grpc"
	"cowait/core"
	"cowait/core/executor"
	"fmt"
	"os"

	"context"

	"go.uber.org/fx"
)

// container client / process manager

func main() {
	executor := fx.New(
		grpc.Module,
		executor.Module,

		fx.Invoke(run),

		// fx.NopLogger,
	)
	executor.Run()
}

func run(exec executor.Executor) {
	// read task data from environment
	id := core.TaskID(os.Getenv(core.EnvTaskID))
	def, err := core.TaskDefFromEnv(os.Getenv(core.EnvTaskdef))
	if err != nil {
		fmt.Println("failed to parse task spec:", err)
		os.Exit(1)
	}

	ctx := context.Background()
	if err := exec.Run(ctx, id, def); err != nil {
		fmt.Println("execution failed:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
