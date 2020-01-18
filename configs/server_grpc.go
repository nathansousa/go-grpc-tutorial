package configs

import (
	"fmt"
	"github.com/nathansousa/go-grpc-example/internal/controllers"
	GRPC "github.com/nathansousa/go-grpc-example/internal/gRPC"
	"google.golang.org/grpc"
	"log"
	"net"
)

func SetUpGRPCServer() {
	getNetListener(8000)
	server := grpc.NewServer()

	controller := controllers.NewBookController()
	GRPC.RegisterBookServiceServer(server, controller)

	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func getNetListener(port uint) net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to get listener", err)
	}

	return listener
}
