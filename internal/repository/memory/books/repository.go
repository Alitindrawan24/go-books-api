package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
	"github.com/Alitindrawan24/go-books-api/internal/schema"
)

type RepositoryProvider interface {
	FindBooks(context.Context, schema.QueryParams) ([]entity.Book, error)
	FindBookByID(context.Context, int) (entity.Book, error)
	InsertBook(context.Context, entity.Book) (entity.Book, error)
	UpdateBook(context.Context, int, entity.Book) (entity.Book, error)
	DeleteBook(context.Context, int) error
}

type Repository struct {
}

func New() RepositoryProvider {
	return &Repository{}
}
