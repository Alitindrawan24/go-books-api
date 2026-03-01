package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
)

var books []entity.Book

func (repository *Repository) FindBooks(ctx context.Context) ([]entity.Book, error) {
	return books, nil
}

func (repository *Repository) FindBookByID(ctx context.Context, id int) (entity.Book, error) {
	book := entity.Book{}
	for _, b := range books {
		if b.ID == id {
			book = b
			break
		}
	}

	return book, nil
}

func (repository *Repository) InsertBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	book.ID = len(books) + 1

	books = append(books, book)

	return book, nil
}
