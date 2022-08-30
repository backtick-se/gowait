package cluster

import (
	"github.com/backtick-se/gowait/core/task"
	"github.com/backtick-se/gowait/util/spy"

	"context"
)

type DriverMock struct {
	KillSpy  spy.E
	PokeSpy  spy.E
	SpawnSpy spy.E
}

var _ Driver = &DriverMock{}

// Kill implements Driver
func (d *DriverMock) Kill(ctx context.Context, id task.ID) error {
	return d.KillSpy.Call(id)
}

// Poke implements Driver
func (d *DriverMock) Poke(ctx context.Context, id task.ID) error {
	return d.PokeSpy.Call(id)
}

// Spawn implements Driver
func (d *DriverMock) Spawn(ctx context.Context, id task.ID, image string) error {
	return d.SpawnSpy.Call(id, image)
}
