package cloud

import (
	"github.com/backtick-se/gowait/core/cluster"

	"context"
	"fmt"

	"go.uber.org/fx"
)

type clustermgr struct {
	server   cluster.UplinkServer
	clusters map[string]*kluster
}

type kluster struct {
	*cluster.Info
	cluster.Client
}

func NewClusterManager(lc fx.Lifecycle, server cluster.UplinkServer) *clustermgr {
	mgr := &clustermgr{
		server:   server,
		clusters: make(map[string]*kluster),
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := mgr.server.Serve(mgr.handle)
				if err != nil {
					fmt.Println("uplink server failed:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return mgr.server.Close()
		},
	})

	return mgr
}

func (s *clustermgr) handle(client cluster.Client) error {
	// query cluster information
	info, err := client.Info(context.Background())
	if err != nil {
		return fmt.Errorf("cluster.Info() failed: %w", err)
	}

	// todo: match against database & existing connections
	// todo: check api key

	cluster := &kluster{
		Info:   info,
		Client: client,
	}

	s.add(cluster)
	defer s.remove(cluster)

	events, err := cluster.Subscribe(context.Background())
	if err != nil {
		return fmt.Errorf("cluster.Subscribe() failed: %w", err)
	}

	for {
		event, ok := <-events.Next()
		if !ok {
			break
		}
		fmt.Println(cluster.Name, "event:", event)
	}

	return nil
}

func (s *clustermgr) add(cluster *kluster) {
	fmt.Println("added cluster:", cluster.Name)
	s.clusters[cluster.ID] = cluster
}

func (s *clustermgr) remove(cluster *kluster) {
	fmt.Println("lost cluster:", cluster.Name)
	delete(s.clusters, cluster.ID)
}
