package main

import (
	"github.com/backtick-se/gowait/adapter/api/grpc"
	"github.com/backtick-se/gowait/core/cli/cmd"

	"fmt"
	"os"

	"go.uber.org/fx"
)

func main() {
	executor := fx.New(
		grpc.Module,
		cmd.Module,

		fx.Invoke(run),

		// fx.NopLogger,
	)
	executor.Run()
}

func run(app cmd.App, shut fx.Shutdowner) error {
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
		return err
	}
	return shut.Shutdown()
}
