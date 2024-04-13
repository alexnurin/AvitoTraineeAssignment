package api

import (
	"github.com/alexnurin/AvitoTraineeAssignment/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"net/http"
)

func getUserBannerHandler(c *gin.Context, db *sqlx.DB) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер для пользователя"})
}

func getAllBannersHandler(c *gin.Context, db *sqlx.DB) {
	var banners []models.Banner
	query := "SELECT * FROM banners"
	if err := db.Select(&banners, query); err != nil {
		log.Printf("ошибка при запросе баннеров: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе баннеров"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"banners": banners})
}

func createBannerHandler(c *gin.Context, db *sqlx.DB) {
	var newBanner models.Banner

	if err := c.ShouldBindJSON(&newBanner); err != nil {
		log.Printf("ошибка при разборе данных баннера: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	var bannerID int
	query := `INSERT INTO banners (tag_ids, feature_id, content, is_active) VALUES ($1, $2, $3, $4) RETURNING banner_id`
	err := db.QueryRow(query, pq.Array(newBanner.TagIDs), newBanner.FeatureID, newBanner.Content, newBanner.IsActive).Scan(&bannerID)
	if err != nil {
		log.Printf("ошибка при добавлении баннера в базу данных: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"banner_id": bannerID})
}

func updateBannerHandler(c *gin.Context, db *sqlx.DB) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер обновлен"})
}

func deleteBannerHandler(c *gin.Context, db *sqlx.DB) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер удален"})
}
