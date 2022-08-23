package daemon

import "go.uber.org/fx"

var Module = fx.Module(
	"daemon",
	fx.Provide(NewDaemon),

	// register daemon.Cluster as the executor handler
	fx.Provide(registerExecutorHandler),
)
