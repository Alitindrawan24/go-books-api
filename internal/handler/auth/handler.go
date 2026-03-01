package auth

import (
	"github.com/Alitindrawan24/go-books-api/internal/usecase/auth"
)

type Handler struct {
	auth auth.UseCaseProvider
}

func New(auth auth.UseCaseProvider) *Handler {
	return &Handler{
		auth: auth,
	}
}
