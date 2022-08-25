package executor

import (
	"github.com/backtick-se/gowait/core/task"

	"context"
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

func run(lc fx.Lifecycle, exec T) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// read task data from environment
			id := task.ID(os.Getenv(task.EnvTaskID))
			def, err := task.SpecFromEnv(os.Getenv(task.EnvTaskdef))
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
