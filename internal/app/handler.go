package app

import "github.com/Alitindrawan24/go-books-api/internal/handler/books"

type Handlers struct {
	BookHandler *books.Handler
}

func NewHandler(useCase *UseCases) *Handlers {
	return &Handlers{
		BookHandler: books.New(useCase.books),
	}
}
