package daemon

import (
	"cowait/core"

	"context"

	"go.uber.org/fx"
)

type uplinkmgr struct {
	uplink core.UplinkClient
}

func NewUplinkManager(lc fx.Lifecycle, uplink core.UplinkClient) *uplinkmgr {
	mgr := &uplinkmgr{
		uplink: uplink,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go mgr.proc()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return uplink.Close()
		},
	})
	return mgr
}

func (u *uplinkmgr) proc() {
	for {
		if err := u.uplink.Connect("cloud.default.svc.cluster.local:1338"); err != nil {
			return
		}
	}
}
