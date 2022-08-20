package client

import "go.uber.org/fx"

var Module = fx.Module(
	"core/client",
	fx.Provide(NewCowaitClient),
	fx.Provide(NewTaskClient),
)
