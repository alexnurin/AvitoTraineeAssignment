package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"time"
)

var globalCache = cache.New(5*time.Minute, 10*time.Minute)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CacheMiddleware())
	return router
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
