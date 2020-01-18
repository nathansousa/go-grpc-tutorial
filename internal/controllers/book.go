package controllers

import (
	"context"
	GRPC "github.com/nathansousa/go-grpc-example/internal/gRPC"
	"log"
)

type BookI interface {
	AddBook(ctx context.Context, book *GRPC.Book) (*GRPC.AddBookResponse, error)
}

type Book struct {
}

func NewBookController() BookI {
	return &Book{}
}

func (b *Book) AddBook(ctx context.Context, book *GRPC.Book) (*GRPC.AddBookResponse, error) {
	log.Println("a new book was inserted")

	return &GRPC.AddBookResponse{
		Book:  book,
		Error: nil,
	}, nil
}

func GetBookAndQuantity(ctx context.Context) (*grpc.GetBookResponse, error) {
	book := &grpc.Book{
		BookID:    "test",
		Name:      "test",
		Author:    "test",
		IsDigital: false,
	}

	return &grpc.GetBookResponse{
		Book:     book,
		Quantity: 10,
		Error:    nil,
	}, nil
}
