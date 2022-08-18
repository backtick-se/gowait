package grpc

import "go.uber.org/fx"

var Module = fx.Module(
	"daemon/grpc",
	fx.Provide(NewTaskServer),

	fx.Invoke(NewServer),
)
