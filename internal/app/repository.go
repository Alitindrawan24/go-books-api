package app

import (
	"github.com/Alitindrawan24/go-books-api/internal/repository/memory/auth"
	"github.com/Alitindrawan24/go-books-api/internal/repository/memory/books"
)

type Repositories struct {
	books books.RepositoryProvider
	auth  auth.RepositoryProvider
}

func NewRepository() *Repositories {
	return &Repositories{
		books: books.New(),
		auth:  auth.New(),
	}
}
