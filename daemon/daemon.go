package daemon

import "go.uber.org/fx"

func New(opts ...fx.Option) *fx.App {
	modules := []fx.Option{}
	modules = append(modules, opts...)
	return fx.New(modules...)
}
