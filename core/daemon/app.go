package daemon

import "go.uber.org/fx"

func App(opts ...fx.Option) *fx.App {
	modules := []fx.Option{
		Module,
	}
	modules = append(modules, opts...)
	return fx.New(modules...)
}
