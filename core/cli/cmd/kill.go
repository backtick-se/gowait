package cmd

import (
	"github.com/backtick-se/gowait/core/api"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"

	"github.com/urfave/cli/v2"
)

type KillCommand cli.ActionFunc

func NewKillCommand(client api.Client) KillCommand {
	return func(cctx *cli.Context) error {
		if cctx.NArg() < 1 {
			return fmt.Errorf("usage: cowait kill <task name>")
		}
		id := cctx.Args().Get(0)

		ctx := context.Background()
		if err := client.KillTask(ctx, task.ID(id)); err != nil {
			return err
		}

		fmt.Println("killed task", id)
		return nil
	}
}
