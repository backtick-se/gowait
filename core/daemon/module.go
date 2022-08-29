package daemon

import "go.uber.org/fx"

var Module = fx.Module(
	"daemon",
	fx.Provide(NewDaemon),
	fx.Provide(NewWorkers),
	fx.Provide(NewTaskManager),

	// register daemon.Cluster as the executor handler
	fx.Provide(registerExecutorHandler),
)
