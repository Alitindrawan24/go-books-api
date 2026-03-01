package auth

import (
	"github.com/gin-gonic/gin"
)

func (handler *Handler) HandleGetToken(c *gin.Context) {
	ctx := c.Request.Context()

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if body.Username != "admin" || body.Password != "password" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	token, err := handler.auth.GetToken(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
