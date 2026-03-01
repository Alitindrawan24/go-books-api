package auth

import (
	"context"

	"github.com/Alitindrawan24/go-books-api/internal/pkg/config"
)

func (repos *Repository) GetToken(ctx context.Context) (string, error) {
	token := config.Get("STATIC_TOKEN")

	return token, nil
}
