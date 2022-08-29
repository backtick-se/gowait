package cmd

import (
	"github.com/urfave/cli/v2"
)

type App interface {
	Run(arguments []string) error
}

type app struct {
	*cli.App
}

func NewApp(
	runAction RunCommand,
	killAction KillCommand,
) App {
	return &app{
		App: &cli.App{
			Name:  "cowait",
			Usage: "cowait helps you run shit",
			Commands: []*cli.Command{
				{
					Name:   "run",
					Usage:  "run a task",
					Action: cli.ActionFunc(runAction),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "image",
							Value: "",
							Usage: "task image name",
						},
					},
				},
				{
					Name:   "kill",
					Usage:  "kill a running task",
					Action: cli.ActionFunc(killAction),
				},
			},
		},
	}
}
