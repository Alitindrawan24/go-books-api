package auth

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/repository/memory/auth"
)

type UseCaseProvider interface {
	GetToken(context.Context) (string, error)
}

type UseCase struct {
	authRepository auth.RepositoryProvider
}

func New(authRepository auth.RepositoryProvider) UseCaseProvider {
	return &UseCase{
		authRepository: authRepository,
	}
}
