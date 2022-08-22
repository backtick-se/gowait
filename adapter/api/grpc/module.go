package grpc

import "go.uber.org/fx"

var Module = fx.Module(
	"daemon/api/grpc",

	fx.Provide(NewServer),

	fx.Provide(NewExecutorClient),
	fx.Provide(NewApiClient),

	fx.Provide(NewUplinkClient),
	fx.Provide(NewUplinkServer),
)
