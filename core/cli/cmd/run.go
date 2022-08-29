package cmd

import (
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"

	"github.com/urfave/cli/v2"
)

type RunCommand cli.ActionFunc

func NewRunCommand(client core.APIClient) RunCommand {
	return func(cctx *cli.Context) error {
		if cctx.NArg() < 1 {
			return fmt.Errorf("usage: cowait run <task name>")
		}
		name := cctx.Args().Get(0)
		image := cctx.String("image")

		ctx := context.Background()
		id, err := client.CreateTask(ctx, &task.Spec{
			Name:    name,
			Image:   image,
			Command: []string{"python", "-um", "cowait"},
		})
		if err != nil {
			return err
		}

		fmt.Println("created task", id)
		return nil
	}
}
