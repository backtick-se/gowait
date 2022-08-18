package daemon

import "go.uber.org/fx"

var Module = fx.Module(
	"daemon",
	fx.Provide(NewTaskManager),
)
