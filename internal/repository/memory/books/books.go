package books

import (
	"context"
	"errors"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
	"github.com/Alitindrawan24/go-books-api/internal/schema"
)

var books []entity.Book

func (repository *Repository) FindBooks(ctx context.Context, queryParams schema.QueryParams) ([]entity.Book, error) {

	// clone books
	filtered := make([]entity.Book, 0, len(books))

	// filter author
	if queryParams.Author != "" {
		for _, b := range books {
			if b.Author == queryParams.Author {
				filtered = append(filtered, b)
			}
		}
	} else {
		filtered = append(filtered, books...)
	}

	total := len(filtered)

	page := queryParams.Page
	limit := queryParams.Limit

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	if offset >= total {
		return []entity.Book{}, nil
	}

	end := offset + limit
	if end > total {
		end = total
	}

	return filtered[offset:end], nil
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
