package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
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

func InitializeRoutes(router *gin.Engine, db *sqlx.DB) {
	adminToken := "admin_token"
	userToken := "user_token"

	anyToken := []string{adminToken, userToken}

	adminRoutes := router.Group("")
	adminRoutes.Use(TokenAuthMiddleware([]string{adminToken}))
	{
		adminRoutes.GET("/banner", func(c *gin.Context) {
			getFilteredBannersHandler(c, db)
		})

		adminRoutes.POST("/banner", func(c *gin.Context) {
			createBannerHandler(c, db)
		})

		adminRoutes.PATCH("/banner/:id", func(c *gin.Context) {
			updateBannerHandler(c, db)
		})

		adminRoutes.DELETE("/banner/:id", func(c *gin.Context) {
			deleteBannerHandler(c, db)
		})
	}
	router.GET("/user_banner", TokenAuthMiddleware(anyToken), func(c *gin.Context) {
		getUserBannerHandler(c, db)
	})
}
