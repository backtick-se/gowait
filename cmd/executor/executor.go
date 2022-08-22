package main

import (
	"cowait/adapter/api/grpc"
	"cowait/core/executor"

	"go.uber.org/fx"
)

// container client / process manager

func main() {
	executor := executor.App(
		grpc.Module,

		fx.Invoke(grpc.RegisterExecutorServer),

		// fx.NopLogger,
	)
	executor.Run()
}
