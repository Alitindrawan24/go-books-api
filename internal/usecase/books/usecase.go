package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
	"github.com/Alitindrawan24/go-books-api/internal/repository/memory/books"
	"github.com/Alitindrawan24/go-books-api/internal/schema"
)

type UseCaseProvider interface {
	GetBooks(context.Context, schema.QueryParams) ([]entity.Book, error)
	GetBook(context.Context, int) (entity.Book, error)
	CreateBook(context.Context, entity.Book) (entity.Book, error)
	UpdateBook(context.Context, int, entity.Book) (entity.Book, error)
	DeleteBook(context.Context, int) error
}

type UseCase struct {
	bookRepository books.RepositoryProvider
}

func New(bookRepository books.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		bookRepository: bookRepository,
	}
}
