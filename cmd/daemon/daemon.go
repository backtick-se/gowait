package main

import (
	"cowait/adapter/api/grpc"
	"cowait/adapter/cluster/k8s"
	"cowait/core"
	"cowait/core/daemon"

	"go.uber.org/fx"
)

func main() {
	cowaitd := daemon.New(
		k8s.Module,

		grpc.Module,

		fx.Invoke(createTask),
	)
	cowaitd.Run()
}

func createTask(mgr daemon.TaskManager) error {
	return mgr.Schedule(&core.TaskDef{
		Name:      "gowait-task",
		Namespace: "default",
		Image:     "cowait/gowait-python",
		Command:   []string{"python", "hello.py"},
	})
}
