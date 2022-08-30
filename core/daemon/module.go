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

	// provide a task queue with a reasonable size
	fx.Provide(func() task.TaskQueue { return task.NewTaskQueue(256) }),

	// register Workers & TaskManager as executor/task RPC handlers
	fx.Provide(func(workers Workers) executor.Handler { return workers }),
	fx.Provide(func(taskMgr task.Manager) task.Handler { return taskMgr }),
)
