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
		fx.NopLogger,
	)
	cowaitd.Run()
}

func createTask(mgr daemon.TaskManager) error {
	_, err := mgr.Schedule(&core.TaskSpec{
		Name:    "gowait-task",
		Image:   "cowait/gowait-python",
		Command: []string{"python", "-u", "hello.py"},
	})
	return err
}
