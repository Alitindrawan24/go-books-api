package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
)

func (u *UseCase) GetBooks(ctx context.Context) ([]entity.Book, error) {
	books, err := u.bookRepository.FindBooks(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *UseCase) GetBook(ctx context.Context, id int) (entity.Book, error) {
	book, err := u.bookRepository.FindBookByID(ctx, id)
	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (u *UseCase) CreateBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	book, err := u.bookRepository.InsertBook(ctx, book)
	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}
