package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
)

type RepositoryProvider interface {
	FindBooks(context.Context) ([]entity.Book, error)
	FindBookByID(ctx context.Context, id int) (entity.Book, error)
	InsertBook(ctx context.Context, book entity.Book) (entity.Book, error)
	UpdateBook(ctx context.Context, id int, book entity.Book) (entity.Book, error)
	DeleteBook(ctx context.Context, id int) error
}

type Repository struct {
}

func New() RepositoryProvider {
	return &Repository{}
}
