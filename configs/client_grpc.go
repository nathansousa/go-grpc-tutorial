package configs

import (
	"google.golang.org/grpc"
	"log"
)

func SetUpGRPCClient() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()
	return conn
}
