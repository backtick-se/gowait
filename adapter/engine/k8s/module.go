package k8s

import "go.uber.org/fx"

var Module = fx.Module(
	"adapter:cluster:kubernetes",
	fx.Provide(NewInCluster),
)
