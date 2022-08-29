package daemon

import (
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"daemon",
	fx.Provide(NewDaemon),
	fx.Provide(NewWorkers),
	fx.Provide(task.NewManager),

	fx.Provide(func(workers Workers) executor.Handler {
		return workers
	}),
	fx.Provide(func(taskMgr task.Manager) task.Handler {
		return taskMgr
	}),
)
