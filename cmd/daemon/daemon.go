package main

import (
	"cowait/adapter/api/grpc"
	"cowait/adapter/engine/k8s"
	"cowait/core"
	"cowait/core/daemon"
	"time"

	"context"

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

func createTask(lc fx.Lifecycle, cluster core.Cluster) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			time.Sleep(3 * time.Second)
			_, err := cluster.Create(context.Background(), &core.TaskSpec{
				Name:    "gowait-task",
				Image:   "cowait/gowait-python",
				Command: []string{"python", "-u", "hello.py"},
			})
			return err
		},
	})
}
