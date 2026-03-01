package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
	"github.com/Alitindrawan24/go-books-api/internal/repository/memory/books"
)

type UseCaseProvider interface {
	GetBooks(context.Context) ([]entity.Book, error)
	GetBook(context.Context, int) (entity.Book, error)
	CreateBook(context.Context, entity.Book) (entity.Book, error)
}

type UseCase struct {
	bookRepository books.RepositoryProvider
}

func New(bookRepository books.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		bookRepository: bookRepository,
	}
}
