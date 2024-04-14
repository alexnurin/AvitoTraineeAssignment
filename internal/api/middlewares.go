package api

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
)

func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.RequestURI()
		if data, found := globalCache.Get(key); found {
			c.Data(http.StatusOK, "application/json; charset=utf-8", data.([]byte))
			c.Abort()
			return
		}

		c.Next()

		if c.Writer.Status() == http.StatusOK {
			responseData, exists := c.Get("response")
			if exists {
				serializedData, ok := responseData.([]byte)
				if ok {
					globalCache.Set(key, serializedData, cache.DefaultExpiration)
				}
			}
		}
	}
}

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
