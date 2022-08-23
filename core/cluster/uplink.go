package cluster

type UplinkClient interface {
	Connect(endpoint string) error
	Close() error
}

type UplinkServer interface {
	Serve(func(Client) error) error
	Close() error
}
