package cmd

import (
	"github.com/backtick-se/gowait/core/api"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"core/cli/cmd",
	fx.Provide(NewApp),

	fx.Provide(NewRunCommand),
	fx.Provide(NewKillCommand),

	fx.Decorate(func(client api.Client) (api.Client, error) {
		hostname := "localhost:1337"
		err := client.Connect(hostname)
		return client, err
	}),
)
