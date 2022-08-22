package grpc

import (
	"cowait/core"

	"fmt"
	"net"

	"github.com/hashicorp/yamux"
)

type uplinkSrv struct {
	listen net.Listener
}

func NewUplinkServer() core.UplinkServer {
	return &uplinkSrv{}
}

func (s *uplinkSrv) Serve(handler func(core.ClusterClient) error) error {
	var err error
	s.listen, err = net.Listen("tcp", fmt.Sprintf(":%d", 1338))
	if err != nil {
		return fmt.Errorf("listen faield: %s", err)
	}

	for {
		// Accept blocks until there is an incoming TCP connection
		incoming, err := s.listen.Accept()
		if err != nil {
			return fmt.Errorf("accept failed: %w", err)
		}

		go s.handle(incoming, handler)
	}
}

func (s *uplinkSrv) Close() error {
	return s.listen.Close()
}

func (u *uplinkSrv) handle(incoming net.Conn, handler func(core.ClusterClient) error) error {
	defer incoming.Close()

	session, err := yamux.Client(incoming, yamux.DefaultConfig())
	if err != nil {
		return fmt.Errorf("create yamux session failed: %w", err)
	}
	defer session.Close()

	muxConn, err := session.Open()
	if err != nil {
		return fmt.Errorf("open yamux connection failed: %w", err)
	}
	defer muxConn.Close()

	fmt.Println("accepted client connection from ", incoming.RemoteAddr())

	client := NewClusterClient()
	if err := client.Connect(muxConn); err != nil {
		return fmt.Errorf("grpc over yamux failed: %w", err)
	}

	if err := handler(client); err != nil {
		return fmt.Errorf("uplink session error: %w", err)
	}

	return nil
}
