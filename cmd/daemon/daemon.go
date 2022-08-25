package main

import (
	"github.com/backtick-se/gowait/adapter/api/grpc"
	"github.com/backtick-se/gowait/adapter/engine/k8s"
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/daemon"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"time"

	"go.uber.org/fx"
)

func main() {
	cowaitd := daemon.App(
		k8s.Module,
		grpc.Module,

		fx.Invoke(grpc.RegisterApiServer),
		fx.Invoke(grpc.RegisterExecutorServer),
		fx.Invoke(daemon.NewUplinkManager),
		fx.Invoke(createTask),
		// fx.NopLogger,
	)
	cowaitd.Run()
}

func createTask(lc fx.Lifecycle, cluster cluster.T) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			time.Sleep(3 * time.Second)
			_, err := cluster.Create(context.Background(), &task.Spec{
				Name:    "gowait-task",
				Image:   "github.com/backtick-se/gowait/gowait-python",
				Command: []string{"python", "-u", "hello.py"},
			})
			return err
		},
	})
}
