package controllers

import (
	"context"
	BookGRPC "github.com/nathansousa/go-grpc-example/internal/gRPC"
	"log"
)

type BookI interface {
	AddBook(ctx context.Context, book *BookGRPC.Book) (*BookGRPC.AddBookResponse, error)
	GetBookAndQuantity(ctx context.Context, req *BookGRPC.GetBookRequest) (*BookGRPC.GetBookResponse, error)
}

type Book struct {
}

func NewBookController() BookI {
	return &Book{}
}

func (b *Book) AddBook(ctx context.Context, book *BookGRPC.Book) (*BookGRPC.AddBookResponse, error) {
	log.Println("a new book was inserted")

	return &BookGRPC.AddBookResponse{
		Book:  book,
		Error: nil,
	}, nil
}

func (b *Book) GetBookAndQuantity(ctx context.Context, req *BookGRPC.GetBookRequest) (*BookGRPC.GetBookResponse, error) {
	log.Println("getting a book")

	book := &BookGRPC.Book{
		BookID:    "test",
		Name:      req.Name,
		Author:    "test",
		IsDigital: false,
	}

	return &BookGRPC.GetBookResponse{
		Book:     book,
		Quantity: 10,
		Error:    nil,
	}, nil
}
