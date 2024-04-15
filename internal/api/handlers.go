package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func getUserBannerHandler(c *gin.Context, db *sqlx.DB) {
	featureID := c.Query("feature_id")
	tagID := c.Query("tag_id")
	useLastRevision := c.Query("use_last_revision") == "true"

	if featureID == "" || tagID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Необходимы параметры feature_id и tag_id"})
		return
	}
	var content json.RawMessage
	var err error

	if useLastRevision {
		query := "SELECT content FROM banners WHERE feature_id = $1 AND $2 = ANY(tag_ids)"
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			query += " AND is_active = true"
		}

		err = db.Get(&content, query, featureID, tagID)
	} else {
		cacheKey := fmt.Sprintf("banner-%s-%s", featureID, tagID)
		cachedContent, found := globalCache.Get(cacheKey)
		if found {
			content = cachedContent.(json.RawMessage)
		} else {
			err = db.Get(&content, "SELECT content FROM banners WHERE feature_id = $1 AND $2 = ANY(tag_ids) ORDER BY updated_at DESC LIMIT 1", featureID, tagID)
			if err == nil {
				globalCache.Set(cacheKey, content, time.Minute*5)
			}
		}
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Баннер не найден"})
		} else {
			log.Printf("ошибка при получении баннера: %s\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		}
		return
	}
	c.JSON(http.StatusOK, content)
}

func getFilteredBannersHandler(c *gin.Context, db *sqlx.DB) {
	params := map[string]string{
		"feature_id": c.Query("feature_id"),
		"tag_id":     c.Query("tag_id"),
		"limit":      c.Query("limit"),
		"offset":     c.Query("offset"),
	}

	query, args := buildGetBannerQuery(params)

	var banners []models.Banner
	if err := db.Select(&banners, query, args...); err != nil {
		log.Printf("ошибка при запросе баннеров: %s\n\t очередь = %s", err, query)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе баннеров"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"banners": banners})
}

func checkBannerUniqueness(db *sqlx.DB, featureID int, tagIDs pq.Int64Array) (int, error) {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM banners WHERE feature_id = $1 AND $2 && tag_ids)`
	err := db.QueryRow(checkQuery, featureID, pq.Array(tagIDs)).Scan(&exists)
	if err != nil {
		log.Printf("ошибка при проверке уникальности баннера: %s\n", err)
		return http.StatusInternalServerError, fmt.Errorf("внутренняя ошибка сервера")
	}
	if exists {
		return http.StatusBadRequest, fmt.Errorf("такая комбинация feature_id и tag_ids уже существует")
	}
	return http.StatusOK, nil
}

func createBannerHandler(c *gin.Context, db *sqlx.DB) {
	var newBanner models.Banner

	if err := c.ShouldBindJSON(&newBanner); err != nil {
		log.Printf("ошибка при разборе данных баннера: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if statusCode, err := checkBannerUniqueness(db, newBanner.FeatureID, newBanner.TagIDs); err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
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

func getCurrentBannerDetails(db *sqlx.DB, bannerID string) (int, pq.Int64Array, error) {
	var featureID int
	var tagIDs pq.Int64Array
	selectQuery := `SELECT feature_id, tag_ids FROM banners WHERE banner_id = $1`
	if err := db.QueryRow(selectQuery, bannerID).Scan(&featureID, pq.Array(&tagIDs)); err != nil {
		return 0, nil, fmt.Errorf("ошибка при получении текущих значений баннера: %w", err)
	}
	return featureID, tagIDs, nil
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
	if input.FeatureID != nil || input.TagIDs != nil {
		currentFeatureID, currentTagIDs, err := getCurrentBannerDetails(db, bannerId)
		if err != nil {
			log.Printf("%v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе текущих данных баннера"})
			return
		}
		if input.FeatureID == nil {
			input.FeatureID = &currentFeatureID
		}
		if input.TagIDs == nil {
			input.TagIDs = &currentTagIDs
		}
		//if statusCode, err := checkBannerUniqueness(db, *input.FeatureID, *input.TagIDs); err != nil {
		//	c.JSON(statusCode, gin.H{"error": err.Error()})
		//	return
		//}
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
