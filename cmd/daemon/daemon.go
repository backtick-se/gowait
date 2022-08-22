package main

import (
	"cowait/adapter/api/grpc"
	"cowait/adapter/engine/k8s"
	"cowait/core"
	"cowait/core/daemon"

	"context"

	"go.uber.org/fx"
)

func main() {
	cowaitd := daemon.New(
		k8s.Module,

		grpc.Module,

		fx.Invoke(createTask),
		// fx.NopLogger,
	)
	cowaitd.Run()
}

func createTask(cluster core.Cluster) error {
	_, err := cluster.Create(context.Background(), &core.TaskSpec{
		Name:    "gowait-task",
		Image:   "cowait/gowait-python",
		Command: []string{"python", "-u", "hello.py"},
	})
	return err
}
