package api

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestGetUserBannerHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("не удалось создать мок базы данных: %s", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			t.Errorf("ошибка при закрытии мока базы данных: %s", err)
		}
	}()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	router := gin.New()
	router.GET("/user_banner", func(c *gin.Context) {
		getUserBannerHandler(c, sqlxDB)
	})

	content := json.RawMessage(`{"title": "test_title", "text": "test_text", "url": "test_url"}`)
	rows := sqlmock.NewRows([]string{"content"}).AddRow(content)

	query := "SELECT content FROM banners WHERE feature_id = $1 AND $2 = ANY(tag_ids) AND is_active = true"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("123", "456").WillReturnRows(rows)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user_banner?feature_id=123&tag_id=456", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "test_title")
	assert.Contains(t, w.Body.String(), "test_text")
	assert.Contains(t, w.Body.String(), "test_url")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все ожидания базы данных были выполнены: %s", err)
	}
	mock.ExpectClose()
}
