package executor

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"executor",
	fx.Provide(NewServer),
	fx.Provide(New),

	// register Server as the executor handler
	fx.Provide(registerExecutorHandler),
)
