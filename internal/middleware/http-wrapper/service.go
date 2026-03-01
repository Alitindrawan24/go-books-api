package httpwrapper

import (
	"context"
	"strings"

	"github.com/Alitindrawan24/go-books-api/internal/constant"
	"github.com/Alitindrawan24/go-books-api/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

// HttpWrapper wraps handlers for HTTP service.
func (middleware *Middleware) HttpWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(string(constant.RequestIDHeader))
		ctx := logger.InitLogCtx(c.Request.Context(), requestID)

		// Client IP
		ipFromHeader := c.GetHeader(string(constant.ClientCFIPAddress))
		clientIP := ipFromHeader

		if clientIP == "" {
			clientIP = c.ClientIP()
			if clientIP == "::1" {
				clientIP = "127.0.0.1"
			}
		}

		clientIP = strings.ReplaceAll(clientIP, ",", "")

		// Host
		host := c.Request.Host

		ctx = context.WithValue(ctx, constant.HostKey, host)
		ctx = context.WithValue(ctx, constant.ClientIPAddress, clientIP)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
