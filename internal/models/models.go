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
