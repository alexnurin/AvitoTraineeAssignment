package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware(validTokens []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("isAdmin", false)
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
			return
		}
		valid := false
		for _, t := range validTokens {
			if token == t {
				valid = true
				break
			}
		}
		if !valid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Пользователь не имеет доступа"})
			return
		}
		c.Set("isAdmin", true)
		c.Next()
	}
}
