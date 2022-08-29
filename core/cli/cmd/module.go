package cmd

import (
	"github.com/backtick-se/gowait/core"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"core/cli/cmd",
	fx.Provide(NewApp),

	fx.Provide(NewRunCommand),
	fx.Provide(NewKillCommand),

	fx.Decorate(func(client core.APIClient) (core.APIClient, error) {
		hostname := "localhost:1337"
		err := client.Connect(hostname)
		return client, err
	}),
)
