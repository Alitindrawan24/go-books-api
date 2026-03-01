package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Alitindrawan24/go-books-api/internal/app"
	"github.com/Alitindrawan24/go-books-api/internal/pkg/config"
	"github.com/Alitindrawan24/go-books-api/internal/pkg/logger"

	"github.com/gosimple/slug"

	_ "github.com/joho/godotenv/autoload"
)

const (
	timeoutServer = 60
)

var (
	port = config.GetWIthDefault("APP_PORT", "8080")
)

type Server struct {
	handler    *app.Handlers
	http       *http.Server
	middleware *app.Middleware
}

func NewHTTP(ctx context.Context) *Server {

	repository := app.NewRepository()
	useCase := app.NewUseCase(repository)
	handler := app.NewHandler(useCase)

	middleware := app.NewMiddleware(useCase)

	logger.Init(slug.Make(config.Get("APP_NAME")))
	logger.Info(ctx, nil, nil, "Connecting - NewHTTP")

	return &Server{
		handler:    handler,
		middleware: middleware,
	}
}

func (s *Server) Run() *http.Server {
	router := NewRouter(s.handler, s.middleware)

	s.http = &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router.Gin,
		ReadTimeout:  timeoutServer * time.Second,
		WriteTimeout: timeoutServer * time.Second,
	}

	fmt.Printf("Server configured on port %s\n", port)

	return s.http
}
