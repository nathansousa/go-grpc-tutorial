package controllers

import (
	"context"

	"github.com/google/uuid"
	bookGRPC "github.com/nathansousa/go-grpc-example/server/pkg/grpc"
	"github.com/sirupsen/logrus"
)

type IBook interface {
	AddBook(context.Context, *bookGRPC.AddBookRequest) (*bookGRPC.AddBookResponse, error)
}

type Book struct {
	bookGRPC.UnimplementedBookStoreServer
}

func NewBookController() *Book {
	return &Book{}
}

func (b *Book) AddBook(_ context.Context, addBookRequest *bookGRPC.AddBookRequest) (*bookGRPC.AddBookResponse, error) {
	logrus.Info("new book inserted: ", addBookRequest)

	return &bookGRPC.AddBookResponse{
		BookID: uuid.New().String(),
	}, nil
}
