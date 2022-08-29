package docker

import "go.uber.org/fx"

var Module = fx.Module(
	"adapter/driver/docker",
	fx.Provide(New),
)
