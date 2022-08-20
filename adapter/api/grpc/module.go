package grpc

import "go.uber.org/fx"

var Module = fx.Module(
	"daemon/api/grpc",
	fx.Provide(NewTaskServer),
	fx.Provide(NewCowaitServer),

	fx.Invoke(NewServer),
)
