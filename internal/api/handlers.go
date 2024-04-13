package api

import (
	"encoding/json"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"net/http"
)

func getUserBannerHandler(c *gin.Context, db *sqlx.DB) {
	featureID := c.Query("feature_id")
	tagID := c.Query("tag_id")
	if featureID == "" || tagID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Необходимы параметры feature_id и tag_id"})
		return
	}
	var contents []json.RawMessage
	query := "SELECT content FROM banners WHERE feature_id = $1 AND $2 = ANY(tag_ids) AND is_active = true"

	err := db.Select(&contents, query, featureID, tagID)
	if err != nil {
		log.Printf("ошибка при получения баннера по feature_id и tag_id: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}
	if len(contents) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Баннеры не найдены"})
		return
	}
	c.JSON(http.StatusOK, contents[0])
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
	bannerId := c.Param("id")
	var input struct {
		TagIDs    *pq.Int64Array   `json:"tag_ids"`
		FeatureID *int             `json:"feature_id"`
		Content   *json.RawMessage `json:"content"`
		IsActive  *bool            `json:"is_active"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}
	query, args := buildUpdateBannerQuery(input, bannerId)
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Printf("ошибка при обновлении баннера: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}
	if count, _ := result.RowsAffected(); count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Баннер не найден"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Баннер успешно обновлен"})
}

func deleteBannerHandler(c *gin.Context, db *sqlx.DB) {
	bannerId := c.Param("id")

	query := `DELETE FROM banners WHERE banner_id = $1`
	result, err := db.Exec(query, bannerId)
	if err != nil {
		log.Printf("Ошибка при удалении баннера: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	if count, _ := result.RowsAffected(); count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Баннер не найден"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
