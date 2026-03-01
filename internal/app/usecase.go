package app

import (
	"github.com/Alitindrawan24/go-books-api/internal/usecase/auth"
	"github.com/Alitindrawan24/go-books-api/internal/usecase/books"
)

type UseCases struct {
	books books.UseCaseProvider
	auth  auth.UseCaseProvider
}

func NewUseCase(repository *Repositories) *UseCases {
	return &UseCases{
		books: books.New(repository.books),
		auth:  auth.New(repository.auth),
	}
}
