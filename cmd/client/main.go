package main

import (
	"context"
	"github.com/nathansousa/go-grpc-example/configs"
	grpc "github.com/nathansousa/go-grpc-example/internal/gRPC"
	"log"
)

func main() {
	client := grpc.NewBookServiceClient(configs.SetUpGRPCClient())

	book := &grpc.Book{
		BookID:    "testClient",
		Name:      "testClient",
		Author:    "testClient",
		IsDigital: true,
	}

	response, err := client.AddBook(context.Background(), book)
	if err != nil {
		log.Println("something wrong happened", err)
	}

	log.Println("book inserted with success:", response)
}
