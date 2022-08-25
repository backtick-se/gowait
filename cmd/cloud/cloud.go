package main

import (
	"github.com/backtick-se/gowait/adapter/api/grpc"
	"github.com/backtick-se/gowait/core/cloud"
)

func main() {
	cloud := cloud.App(
		grpc.Module,
	)
	cloud.Run()
}
