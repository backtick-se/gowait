package main

import (
	"cowait/adapter/api/grpc"
	"cowait/core/cloud"
)

func main() {
	cloud := cloud.App(
		grpc.Module,
	)
	cloud.Run()
}
