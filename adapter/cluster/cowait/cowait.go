package cowait

import (
	"context"
	"cowait/core"
	"cowait/core/client"
)

// cowait clusters recursively use other cowait daemons as clusters
type cluster struct {
	name string
	api  client.Client
}

var _ core.Cluster = &cluster{}

func (c *cluster) Name() string {
	return c.name
}

func (c *cluster) Spawn(ctx context.Context, id core.TaskID, def *core.TaskSpec) error {
	_, err := c.api.CreateTask(ctx, id, def)
	return err
}

func (c *cluster) Kill(context.Context, core.TaskID) error {
	return nil
}

func (c *cluster) Poke(context.Context, core.TaskID) error {
	return nil
}
