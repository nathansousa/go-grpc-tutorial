package main

import (
	"context"

	"github.com/nathansousa/go-grpc-example/client/config"
	bookGRPC "github.com/nathansousa/go-grpc-example/server/pkg/grpc"
	"github.com/sirupsen/logrus"
)

func main() {
	conn := config.SetupGRPCClient("localhost:8000")

	bookStoreClient := bookGRPC.NewBookStoreClient(conn)

	book := &bookGRPC.AddBookRequest{
		Name:   "Grpc tutorial",
		Author: "Nathan de Sousa Martins",
	}

	response, err := bookStoreClient.AddBook(context.Background(), book)
	if err != nil {
		logrus.Errorf("failed to add the new book: %s", err)
	}

	logrus.Info("a new was book added with id: ", response.BookID)
}
