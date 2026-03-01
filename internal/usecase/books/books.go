package books

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
	"github.com/Alitindrawan24/go-books-api/internal/schema"
)

func (u *UseCase) GetBooks(ctx context.Context, queryParams schema.QueryParams) ([]entity.Book, error) {
	books, err := u.bookRepository.FindBooks(ctx, queryParams)
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

func (u *UseCase) UpdateBook(ctx context.Context, id int, book entity.Book) (entity.Book, error) {
	book, err := u.bookRepository.UpdateBook(ctx, id, book)
	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (u *UseCase) DeleteBook(ctx context.Context, id int) error {
	err := u.bookRepository.DeleteBook(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
