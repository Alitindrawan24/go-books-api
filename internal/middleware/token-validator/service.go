package tokenvalidator

import (
	"github.com/Alitindrawan24/go-books-api/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

func (middleware *Middleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		if bearerToken == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		staticToken := config.Get("STATIC_TOKEN")
		if bearerToken != "Bearer "+staticToken {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
