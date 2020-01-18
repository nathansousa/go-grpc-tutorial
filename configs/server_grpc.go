package configs

import (
	"fmt"
	BookGRPC "github.com/nathansousa/go-grpc-example/internal/gRPC"
	"github.com/nathansousa/go-grpc-example/internal/gRPC/controllers"
	"google.golang.org/grpc"
	"log"
	"net"
)

func SetUpGRPCServer() {
	serverGRPC := grpc.NewServer()
	BookGRPC.RegisterBookServiceServer(serverGRPC, controllers.NewBookController())

	err := serverGRPC.Serve(getNetListener(8000))
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Println("gRPC server is running on port: 8000")
}

func getNetListener(port uint) net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to get listener", err)
	}

	return listener
}
