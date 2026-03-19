package middleware

import (
	"main/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Missing Authorization header",
			})
			return
		}

		// Support both "Bearer sk-xxx" and "sk-xxx"
		token := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		if !strings.HasPrefix(token, "sk-") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid API key format (must start with 'sk-')",
			})
			return
		}

		if token != cfg.APIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid API key",
			})
			return
		}

		c.Next()
	}
}
