package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func InitializeRoutes(router *gin.Engine, db *sqlx.DB) {
	router.GET("/user_banner", func(c *gin.Context) {
		getUserBannerHandler(c, db)
	})

	router.GET("/banner", func(c *gin.Context) {
		getAllBannersHandler(c, db)
	})

	router.POST("/banner", func(c *gin.Context) {
		createBannerHandler(c, db)
	})

	router.PATCH("/banner/:id", func(c *gin.Context) {
		updateBannerHandler(c, db)
	})

	router.DELETE("/banner/:id", func(c *gin.Context) {
		deleteBannerHandler(c, db)
	})

}
