package main

import (
	"cowait/core/client"
	"cowait/executor"

	"context"

	"go.uber.org/fx"
)

// container client / process manager

func main() {
	executor := fx.New(
		client.Module,

		fx.Provide(executor.NewFromEnv),
		fx.Invoke(run),

		fx.NopLogger,
	)
	executor.Run()
}

func run(lc fx.Lifecycle, exec executor.Executor) error {
	ctx := context.Background()
	return exec.Run(ctx)
}
