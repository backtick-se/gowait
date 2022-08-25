package cloud

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/task"
	"github.com/backtick-se/gowait/util/events"

	"context"
	"fmt"
)

type kluster struct {
	info   *cluster.Info
	client cluster.Client
	events events.Pub[*cluster.Event]
}

var _ cluster.T = &kluster{}

func (c *kluster) proc() error {
	events, err := c.client.Subscribe(context.Background())
	if err != nil {
		return fmt.Errorf("cluster.Subscribe() failed: %w", err)
	}

	for {
		event, ok := <-events.Next()
		if !ok {
			break
		}
		fmt.Println(c.info.Name, "event:", event)
		c.events.Publish(event)
	}

	return nil
}

func (c *kluster) Get(ctx context.Context, id task.ID) (task.T, bool) {
	return nil, false
}

func (c *kluster) Create(ctx context.Context, spec *task.Spec) (task.T, error) {
	state, err := c.client.CreateTask(ctx, spec)
	if err != nil {
		return nil, err
	}
	return &instance{
		cluster: c,
		state:   state,
	}, nil
}

func (c *kluster) Destroy(ctx context.Context, id task.ID) error {
	return nil
}

func (c *kluster) Events() events.Pub[*cluster.Event] {
	return c.events
}

func (c *kluster) Info() cluster.Info {
	return *c.info
}
