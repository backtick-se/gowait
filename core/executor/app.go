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
			go func() {
				id := task.ID(os.Getenv(task.EnvTaskID))
				ctx := context.Background()
				if err := exec.Run(ctx, id); err != nil {
					fmt.Println("execution failed:", err)
					os.Exit(1)
				}
				os.Exit(0)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}
