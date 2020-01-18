package main

import (
	"context"
	"github.com/nathansousa/go-grpc-example/configs"
	bookGRPC "github.com/nathansousa/go-grpc-example/internal/gRPC"
	"log"
)

func main() {
	client := bookGRPC.NewBookServiceClient(configs.SetUpGRPCClient())

	for i := 0; i < 10; i++ {
		AddBook(string(i), client)
		GetBook(string(i), client)
	}
}

func AddBook(data string, client bookGRPC.BookServiceClient) {
	book := &bookGRPC.Book{
		BookID:    data,
		Name:      data,
		Author:    data,
		IsDigital: true,
	}

	response, err := client.AddBook(context.Background(), book)
	if err != nil {
		log.Println("something wrong happened", err)
	}

	log.Println("book inserted with success:", response)
}

func GetBook(name string, client bookGRPC.BookServiceClient) {
	request := &bookGRPC.GetBookRequest{
		Name: name,
	}

	response, err := client.GetBookAndQuantity(context.Background(), request)
	if err != nil {
		log.Println("something wrong happened", err)
	}

	log.Println("success get book data:", response)
}
