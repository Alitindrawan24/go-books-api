package books

import "github.com/Alitindrawan24/go-books-api/internal/usecase/books"

type Handler struct {
	books books.UseCaseProvider
}

func New(books books.UseCaseProvider) *Handler {
	return &Handler{
		books: books,
	}
}
