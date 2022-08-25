package cloud

import (
	"github.com/backtick-se/gowait/core/cluster"

	"context"
	"fmt"

	"go.uber.org/fx"
)

// Some kind of service to query a defined cluster based on an ID and a secret key
type ClusterFinder interface {
	Find(id, key string) (*cluster.Info, error)
}

type clustermgr struct {
	server   cluster.UplinkServer
	clusters map[string]cluster.T
	finder   ClusterFinder
}

func NewClusterManager(lc fx.Lifecycle, server cluster.UplinkServer) *clustermgr {
	mgr := &clustermgr{
		server:   server,
		clusters: make(map[string]cluster.T),
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
	actual, err := s.finder.Find(info.ID, info.Key)
	if err != nil {
		return err
	}

	cluster := &kluster{
		info:   actual,
		client: client,
	}

	s.add(cluster)
	defer s.remove(cluster)

	return cluster.proc()
}

func (s *clustermgr) add(cluster *kluster) {
	fmt.Println("added cluster:", cluster.info.Name)
	s.clusters[cluster.info.ID] = cluster
}

func (s *clustermgr) remove(cluster *kluster) {
	fmt.Println("lost cluster:", cluster.info.Name)
	delete(s.clusters, cluster.info.ID)
}
