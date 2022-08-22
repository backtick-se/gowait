package core

type UplinkClient interface {
	Connect(endpoint string) error
	Close() error
}

type UplinkServer interface {
	Serve(func(ClusterClient) error) error
	Close() error
}
