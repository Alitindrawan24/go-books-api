package app

import (
	httpwrapper "github.com/Alitindrawan24/go-books-api/internal/middleware/http-wrapper"
)

// Middleware types of middleware layer.
type Middleware struct {
	HttpWrapper *httpwrapper.Middleware
}

// NewMiddleware initializes middleware
func NewMiddleware(useCase *UseCases) *Middleware {
	return &Middleware{
		HttpWrapper: httpwrapper.New(),
	}
}
