package main

import (
	"github.com/backtick-se/gowait/adapter/api/grpc"
	"github.com/backtick-se/gowait/core/executor"

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
