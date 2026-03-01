package books

import (
	"context"
	"errors"

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

	if book.ID == 0 {
		return entity.Book{}, errors.New("book not found")
	}

	return book, nil
}

func (repository *Repository) InsertBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	book.ID = len(books) + 1

	books = append(books, book)

	return book, nil
}

func (repository *Repository) UpdateBook(ctx context.Context, id int, book entity.Book) (entity.Book, error) {
	for i, b := range books {
		if b.ID == id {
			books[i] = book
			books[i].ID = id
			return books[i], nil
		}
	}

	// handle not found
	if book.ID == 0 {
		return entity.Book{}, errors.New("book not found")
	}

	return book, nil
}

func (repository *Repository) DeleteBook(ctx context.Context, id int) error {
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}

	return nil
}
