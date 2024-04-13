package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func InitializeRoutes(router *gin.Engine) {
	router.GET("/user_banner", getUserBannerHandler)
	router.GET("/banner", getAllBannersHandler)
	router.POST("/banner", createBannerHandler)
	router.PATCH("/banner/:id", updateBannerHandler)
	router.DELETE("/banner/:id", deleteBannerHandler)
}
