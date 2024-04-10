package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUserBannerHandler(c *gin.Context) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер для пользователя"})
}

func getAllBannersHandler(c *gin.Context) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Все баннеры"})
}

func createBannerHandler(c *gin.Context) {
	// TODO business logic
	c.JSON(http.StatusCreated, gin.H{"message": "Баннер создан"})
}

func updateBannerHandler(c *gin.Context) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер обновлен"})
}

func deleteBannerHandler(c *gin.Context) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер удален"})
}
