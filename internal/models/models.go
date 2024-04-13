package models

import (
	"encoding/json"
	"github.com/lib/pq"
	"time"
)

type Banner struct {
	BannerID  int             `json:"banner_id" db:"banner_id"`
	TagIDs    pq.Int64Array   `json:"tag_ids,omitempty" db:"tag_ids"`
	FeatureID int             `json:"feature_id" db:"feature_id"`
	Content   json.RawMessage `json:"content,omitempty" db:"content"`
	IsActive  bool            `json:"is_active" db:"is_active"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
}

type UserBannerRequest struct {
	TagID           int  `form:"tag_id" binding:"required"`
	FeatureID       int  `form:"feature_id" binding:"required"`
	UseLastRevision bool `form:"use_last_revision"`
}

type BannerFilterRequest struct {
	FeatureID int `form:"feature_id"`
	TagID     int `form:"tag_id"`
	Limit     int `form:"limit"`
	Offset    int `form:"offset"`
}

type CreateBannerRequest struct {
	TagIDs    []int                  `json:"tag_ids" binding:"required"`
	FeatureID int                    `json:"feature_id" binding:"required"`
	Content   map[string]interface{} `json:"content" binding:"required"`
	IsActive  bool                   `json:"is_active"`
}

type UpdateBannerRequest struct {
	TagIDs    []int                  `json:"tag_ids"`
	FeatureID int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}
