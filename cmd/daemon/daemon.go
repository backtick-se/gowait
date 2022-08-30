package main

import (
	"github.com/backtick-se/gowait/adapter/api/grpc"
	"github.com/backtick-se/gowait/adapter/driver/docker"
	"github.com/backtick-se/gowait/adapter/driver/k8s"
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/daemon"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"go.uber.org/fx"
)

func main() {
	driverPtr := flag.String("driver", "kubernetes", "executor driver id")
	flag.Parse()

	var driver fx.Option
	switch *driverPtr {
	case "kubernetes":
		driver = k8s.Module
	case "docker":
		driver = docker.Module
	default:
		fmt.Println("unknown driver:", *driverPtr)
		os.Exit(1)
	}

	cowaitd := daemon.App(
		driver,
		grpc.Module,

		fx.Invoke(grpc.RegisterApiServer),
		fx.Invoke(grpc.RegisterExecutorServer),
		fx.Invoke(daemon.NewUplinkManager),
		fx.Invoke(createTask),

		fx.NopLogger,
	)
	cowaitd.Run()
}

func createTask(lc fx.Lifecycle, cluster cluster.T) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			time.Sleep(2 * time.Second)
			queueTask := func(name string) {
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				_, err := cluster.Create(ctx, &task.Spec{
					Name:    name,
					Image:   "cowait/gowait-python",
					Command: []string{"python", "-um", "cowait"},
				})
				if err != nil {
					fmt.Println("failed to spawn executor", name, ":", err)
				}
			}

			queueTask("cowait.builtin.enumerate")

			go func() {
				time.Sleep(20 * time.Second)
				queueTask("hello.my_task")
			}()

			return nil
		},
	})
}
