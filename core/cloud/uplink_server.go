package cloud

import (
	"cowait/adapter/api/grpc"
	"cowait/core/client"

	"context"
	"fmt"
	"net"

	"github.com/hashicorp/yamux"
	"go.uber.org/fx"
)

type uplinkSrv struct {
	clusters map[string]client.Cluster
}

func NewUplinkServer(lc fx.Lifecycle) *uplinkSrv {
	srv := &uplinkSrv{}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := srv.Serve()
				if err != nil {
					fmt.Println("uplink server failed:", err)
				}
			}()
			return nil
		},
	})

	return srv
}

func (s *uplinkSrv) Serve() error {
	// start tcp listener on all interfaces
	// note that each connection consumes a file descriptor
	// you may need to increase your fd limits if you have many concurrent clients
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 1338))
	if err != nil {
		return fmt.Errorf("could not listen: %s", err)
	}
	defer ln.Close()

	fmt.Println("waiting for incoming client connections...")
	for {
		// Accept blocks until there is an incoming TCP connection
		incoming, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("couldn't accept %s", err)
		}

		incomingConn, err := yamux.Client(incoming, yamux.DefaultConfig())
		if err != nil {
			return fmt.Errorf("couldn't create yamux %s", err)
		}

		conn, err := incomingConn.Open()
		if err != nil {
			fmt.Println("failed top yamux connection:", err)
			continue
		}
		fmt.Println("accepted client connection from ", incoming.RemoteAddr())

		client := grpc.NewClusterClient()
		if err := client.Connect(conn); err != nil {
			fmt.Println("failed top open grpc over yamux:", err)
			continue
		}

		// handle connection in goroutine so we can accept new TCP connections
		go s.handle(client)
	}
}

func (s *uplinkSrv) handle(client client.Cluster) {
	reply, err := client.Info(context.Background())
	if err != nil {
		fmt.Println("cluster.Info() failed:", err)
		return
	}
	name := reply.Name
	fmt.Println("added cluster:", name)
	s.clusters[name] = client

	events, err := client.Subscribe(context.Background())
	if err != nil {
		fmt.Println("cluster.Subscribe() failed:", err)
	}

	for {
		event, ok := events.Read()
		if !ok {
			break
		}
		fmt.Println(name, "event:", event)
	}

	fmt.Println("lost cluster:", name)
	delete(s.clusters, name)
}
