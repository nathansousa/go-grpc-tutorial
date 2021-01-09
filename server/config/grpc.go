package config

import (
	"fmt"
	"net"

	bookController "github.com/nathansousa/go-grpc-example/server/pkg/controllers"
	bookGRPC "github.com/nathansousa/go-grpc-example/server/pkg/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func SetupGRPCServer(port uint) {
	server := registerGRPCServiceMethods(grpc.NewServer())

	logrus.Info("GRPC server is running on port: ", port)

	if err := server.Serve(getNetListener(port)); err != nil {
		logrus.Panicf("failed to setup grpc server: %s", err)
	}
}

func getNetListener(port uint) net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Panicf("failed to get listener: %s", err)
	}

	return listener
}

func registerGRPCServiceMethods(server *grpc.Server) *grpc.Server {
	bookGRPC.RegisterBookStoreServer(server, bookController.NewBookController())

	return server
}
