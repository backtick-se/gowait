package main

import (
	"context"
	"cowait/adapter/api/grpc"
	"cowait/adapter/cluster/k8s"
	"cowait/core"
	"cowait/core/daemon"
	// "go.uber.org/fx"
)

func main() {
	cowaitd := daemon.New(
		k8s.Module,

		grpc.Module,

		// fx.Invoke(createTask),
	)
	cowaitd.Run()
}

func createTask(cluster core.Cluster) error {
	_, err := cluster.Spawn(context.Background(), &core.TaskDef{
		Name:      "gowait-task",
		Namespace: "default",
		Image:     "cowait/gowait-python",
		Command:   []string{"python", "hello.py"},
	})
	return err
}
