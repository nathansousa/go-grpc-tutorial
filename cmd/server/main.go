package main

import (
	"bitbucket-repository-management-service/internal/grpc/impl"
	"bitbucket-repository-management-service/internal/grpc/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {


	// start the server
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
