package main

import (
	"cowait/core"
	"cowait/core/client"
	"cowait/executor"
	"os"

	"context"

	"go.uber.org/fx"
)

// container client / process manager

func main() {
	executor := fx.New(
		client.Module,
		executor.Module,

		fx.Invoke(run),

		fx.NopLogger,
	)
	executor.Run()
}

func run(lc fx.Lifecycle, exec executor.Executor) error {
	// read task data from environment
	id := core.TaskID(os.Getenv(core.EnvTaskID))
	def, err := core.TaskDefFromEnv(os.Getenv(core.EnvTaskdef))
	if err != nil {
		return err
	}

	ctx := context.Background()
	return exec.Run(ctx, id, def)
}
