package docker

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type dock struct {
	client *client.Client
}

func New() (cluster.Driver, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &dock{
		client: cli,
	}, nil
}

// Kill implements cluster.Driver
func (d *dock) Kill(ctx context.Context, id task.ID) error {
	return d.client.ContainerRemove(ctx, string(id), types.ContainerRemoveOptions{
		Force: true,
	})
}

// Poke implements cluster.Driver
func (d *dock) Poke(ctx context.Context, id task.ID) error {
	r, err := d.client.ContainerInspect(ctx, string(id))
	if err != nil {
		// todo: we need to differentiate between container errors and API errors
		return err
	}
	if r.State.Status != "running" {
		return fmt.Errorf("container state is %s", r.State.Status)
	}
	return nil
}

// Spawn implements cluster.Driver
func (d *dock) Spawn(ctx context.Context, id task.ID, image string) error {
	cfg := container.Config{
		Image: image,
		Env: []string{
			fmt.Sprintf("%s=%s", task.EnvTaskID, id),
			fmt.Sprintf("%s=%s", "COWAIT_IMAGE", image),
			fmt.Sprintf("%s=%s", "COWAIT_DAEMON", "daemon:1337"),
		},
	}
	host := container.HostConfig{
		AutoRemove: true,
	}
	net := network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"cowait": {},
		},
	}
	name := string(id)
	r, err := d.client.ContainerCreate(ctx, &cfg, &host, &net, nil, name)
	if err != nil {
		return err
	}

	return d.client.ContainerStart(ctx, r.ID, types.ContainerStartOptions{})
}
