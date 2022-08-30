package executor

import (
	"github.com/backtick-se/gowait/core/task"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"executor",
	fx.Provide(NewServer),
	fx.Provide(New),

	// register Server as task & executor RPC handler
	fx.Provide(func(server Server) Handler { return server }),
	fx.Provide(func(server Server) task.Handler { return server }),
)
