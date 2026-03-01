package auth

import (
	"context"
)

type RepositoryProvider interface {
	GetToken(context.Context) (string, error)
}

type Repository struct {
}

func New() RepositoryProvider {
	return &Repository{}
}
