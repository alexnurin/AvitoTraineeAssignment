package api

import (
	"fmt"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func getUserBannerHandler(c *gin.Context, db *sqlx.DB) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер для пользователя"})
}

func getAllBannersHandler(c *gin.Context, db *sqlx.DB) {
	var banners []models.Banner
	rows, err := db.Query("SELECT banner_id FROM banners")
	if err != nil {
		fmt.Printf("ошибка при запросе баннеров: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе баннеров"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var banner models.Banner
		if err := rows.Scan(&banner.BannerID, &banner.BannerData, &banner.FeatureID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при чтении данных баннера"})
			return
		}
		banners = append(banners, banner)
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обработке данных"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"banners": banners})
}

func createBannerHandler(c *gin.Context, db *sqlx.DB) {
	// TODO business logic
	c.JSON(http.StatusCreated, gin.H{"message": "Баннер создан"})
}

func updateBannerHandler(c *gin.Context, db *sqlx.DB) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер обновлен"})
}

func deleteBannerHandler(c *gin.Context, db *sqlx.DB) {
	// TODO business logic
	c.JSON(http.StatusOK, gin.H{"message": "Баннер удален"})
}
