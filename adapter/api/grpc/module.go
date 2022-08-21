package grpc

import "go.uber.org/fx"

var Module = fx.Module(
	"daemon/api/grpc",
	fx.Provide(NewExecutorServer),
	fx.Provide(NewExecutorClient),
	fx.Provide(NewApiServer),
	fx.Provide(NewApiClient),

	fx.Invoke(NewServer),
)
