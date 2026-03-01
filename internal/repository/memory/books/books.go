package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
)

func (repository *Repository) FindBooks(ctx context.Context) ([]entity.Book, error) {

	books := []entity.Book{}

	return books, nil
}

func (repository *Repository) FindBookByID(ctx context.Context, id int) (entity.Book, error) {
	return entity.Book{}, nil
}

func (repository *Repository) InsertBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	return entity.Book{}, nil
}
