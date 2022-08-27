package main

import (
	"fmt"

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
			time.Sleep(2 * time.Second)
			queueTask := func(name string) {
				cluster.Create(context.Background(), &task.Spec{
					Name:    name,
					Image:   "cowait/gowait-python",
					Command: []string{"python", "-um", "cowait"},
				})
			}

			queueTask("cowait.builtin.enumerate")

			go func() {
				time.Sleep(20 * time.Second)
				fmt.Println("queueing second task!")
				queueTask("hello.my_task")
			}()

			return nil
		},
	})
}
