package docker

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/task"
	"go.uber.org/zap"

	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type dock struct {
	log    *zap.Logger
	client *client.Client
}

func New(log *zap.Logger) (cluster.Driver, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	log = log.With(zap.String("driver", "docker"))
	log.Info("using docker driver")

	return &dock{
		log:    log,
		client: cli,
	}, nil
}

// Kill implements cluster.Driver
func (d *dock) Kill(ctx context.Context, id task.ID) error {
	d.log.Debug("kill container", zap.String("container_id", string(id)))
	return d.client.ContainerRemove(ctx, string(id), types.ContainerRemoveOptions{
		Force: true,
	})
}

// Poke implements cluster.Driver
func (d *dock) Poke(ctx context.Context, id task.ID) error {
	d.log.Debug("poke container", zap.String("container_id", string(id)))
	r, err := d.client.ContainerInspect(ctx, string(id))
	if err != nil {
		d.log.Error("failed to inspect container", zap.String("container_id", string(id)), zap.Error(err))
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
	d.log.Debug("create container", zap.String("container_id", string(id)))
	cfg := container.Config{
		Image: image,
		Env: []string{
			fmt.Sprintf("%s=%s", task.EnvTaskID, id),
			fmt.Sprintf("%s=%s", "COWAIT_IMAGE", image),
			fmt.Sprintf("%s=%s", "COWAIT_DAEMON", "daemon:1337"),
		},
	}
	host := container.HostConfig{
		AutoRemove: false,
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

	d.log.Debug("start container", zap.String("container_id", string(id)))
	return d.client.ContainerStart(ctx, r.ID, types.ContainerStartOptions{})
}
