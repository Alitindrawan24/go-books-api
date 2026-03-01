package server

import (
	"io"

	app "github.com/Alitindrawan24/go-books-api/internal/app"
	"github.com/Alitindrawan24/go-books-api/internal/pkg/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin        *gin.Engine
	Handler    *app.Handlers
	Middleware *app.Middleware
}

var (
	isLocal       = config.GetWIthDefault("APP_ENV", "development") == "local"
	isDevelopment = config.GetWIthDefault("APP_ENV", "development") == "development"
)

func NewRouter(handler *app.Handlers, middleware *app.Middleware) *Router {
	router := &Router{
		Gin:        gin.Default(),
		Handler:    handler,
		Middleware: middleware,
	}

	router.Gin.Use(middleware.HttpWrapper.HttpWrapper())

	tokenMiddleware := []gin.HandlerFunc{
		middleware.TokenValidator.ValidateToken(),
	}

	RegisterPingRoutes(router)
	RegisterEchoRoutes(router)
	RegisterBookRoutes(router, tokenMiddleware...)
	RegisterAuth(router)

	return router
}

func RegisterPingRoutes(router *Router, m ...gin.HandlerFunc) {
	router.Gin.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
}

func RegisterEchoRoutes(router *Router, m ...gin.HandlerFunc) {
	router.Gin.POST("echo", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid body"})
			return
		}

		c.Data(200, "application/json", body)
	})
}

func RegisterBookRoutes(router *Router, m ...gin.HandlerFunc) {
	router.Gin.GET("books", m[0], router.Handler.BookHandler.HandleGetBooks)
	router.Gin.GET("books/:id", router.Handler.BookHandler.HandleGetBook)
	router.Gin.POST("books", router.Handler.BookHandler.HandleStoreBook)
	router.Gin.PUT("books/:id", router.Handler.BookHandler.HandleUpdateBook)
	router.Gin.DELETE("books/:id", router.Handler.BookHandler.HandleDeleteBook)
}

func RegisterAuth(router *Router, m ...gin.HandlerFunc) {
	router.Gin.POST("auth/token", router.Handler.AuthHandler.HandleGetToken)
}
