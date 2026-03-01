package app

import "github.com/Alitindrawan24/go-books-api/internal/repository/memory/books"

type Repositories struct {
	books books.RepositoryProvider
}

func NewRepository() *Repositories {
	return &Repositories{
		books: books.New(),
	}
}
