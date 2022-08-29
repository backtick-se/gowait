package k8s

import "go.uber.org/fx"

var Module = fx.Module(
	"adapter/driver/kubernetes",
	fx.Provide(NewInCluster),
)
