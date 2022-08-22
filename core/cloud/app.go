package cloud

import "go.uber.org/fx"

var Module = fx.Module(
	"cloud",
	fx.Invoke(NewClusterManager),
)

func App(opts ...fx.Option) *fx.App {
	modules := []fx.Option{
		Module,
	}
	modules = append(modules, opts...)
	return fx.New(modules...)
}
