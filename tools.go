//go:build tools

package main

import (
	_ "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
