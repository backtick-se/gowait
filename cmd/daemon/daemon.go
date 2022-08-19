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
		ID:      "gowait-task-123",
		Name:    "gowait-task",
		Image:   "cowait/gowait-python1233",
		Command: []string{"python", "-u", "hello.py"},
	})
}
