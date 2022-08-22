package executor

import (
	"context"
	"cowait/core"
	"fmt"
	"os"

	"go.uber.org/fx"
)

func App(opts ...fx.Option) *fx.App {
	modules := []fx.Option{
		Module,
	}
	modules = append(modules, opts...)
	modules = append(modules, fx.Invoke(run))
	return fx.New(modules...)
}

func run(lc fx.Lifecycle, exec Executor) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
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
			return nil
		},
	})
}
