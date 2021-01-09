package main

import "github.com/nathansousa/go-grpc-example/server/config"

func main() {
	config.SetupGRPCServer(8000)
}
