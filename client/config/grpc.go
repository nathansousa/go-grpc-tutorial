package config

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func SetupGRPCClient(grpcTarget string) *grpc.ClientConn {
	conn, err := grpc.Dial(grpcTarget, grpc.WithInsecure())
	if err != nil {
		logrus.Panicf("failed to connect to grpc server: %s", err)
	}

	return conn
}
