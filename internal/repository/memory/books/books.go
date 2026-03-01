package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
)

func (repository *Repository) FindBooks(ctx context.Context) ([]entity.Book, error) {

	books := []entity.Book{
		{
			ID:     1,
			Title:  "Book Title",
			Author: "Book Author",
			Year:   2023,
		},
		{
			ID:     2,
			Title:  "Book Title",
			Author: "Book Author",
			Year:   2023,
		},
		{
			ID:     3,
			Title:  "Book Title",
			Author: "Book Author",
			Year:   2023,
		},
	}

	return books, nil
}

func (repository *Repository) FindBookByID(ctx context.Context, id int) (entity.Book, error) {
	return entity.Book{
		ID:     id,
		Title:  "Book Title",
		Author: "Book Author",
		Year:   2023,
	}, nil
}

func (repository *Repository) InsertBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	return entity.Book{
		ID:     1,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}, nil
}
