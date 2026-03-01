package server

import (
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

	RegisterPingRoutes(router)
	RegisterEchoRoutes(router)

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
		var json map[string]interface{}
		c.BindJSON(&json)
		c.JSON(200, json)
	})
}
