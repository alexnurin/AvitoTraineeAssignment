package models

import (
	"encoding/json"
	"time"
)

type Feature struct {
	FeatureID int `db:"feature_id"`
}

type Tag struct {
	TagID int `db:"tag_id"`
}

type Banner struct {
	BannerID   int             `db:"banner_id"`
	BannerData json.RawMessage `db:"banner_data"`
	FeatureID  int             `db:"feature_id"`
}

type BannerTag struct {
	BannerID int `db:"banner_id"`
	TagID    int `db:"tag_id"`
}

type User struct {
	UserID int  `db:"user_id"`
	IsVIP  bool `db:"is_vip"`
}

type UserTag struct {
	UserID int `db:"user_id"`
	TagID  int `db:"tag_id"`
}

type UserBannerVisibility struct {
	UserID     int       `db:"user_id"`
	BannerID   int       `db:"banner_id"`
	IsCurrent  bool      `db:"is_current"`
	LastUpdate time.Time `db:"last_update"`
}
