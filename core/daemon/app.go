package daemon

import (
	"github.com/backtick-se/gowait/util"

	"go.uber.org/fx"
)

func App(opts ...fx.Option) *fx.App {
	modules := []fx.Option{
		Module,
		util.LogModule,
	}
	modules = append(modules, opts...)
	return fx.New(modules...)
}
