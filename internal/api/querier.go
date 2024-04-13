package api

import (
	"encoding/json"
	"github.com/lib/pq"
	"strconv"
	"time"
)

type BannerUpdateInput struct {
	TagIDs    *pq.Int64Array   `json:"tag_ids"`
	FeatureID *int             `json:"feature_id"`
	Content   *json.RawMessage `json:"content"`
	IsActive  *bool            `json:"is_active"`
}

func buildUpdateBannerQuery(input BannerUpdateInput, bannerID string) (string, []interface{}) {
	query := "UPDATE banners SET updated_at = $1"
	args := []interface{}{time.Now()}
	i := 2

	if input.TagIDs != nil {
		query += ", tag_ids = $" + strconv.Itoa(i)
		args = append(args, pq.Array(*input.TagIDs))
		i++
	}
	if input.FeatureID != nil {
		query += ", feature_id = $" + strconv.Itoa(i)
		args = append(args, *input.FeatureID)
		i++
	}
	if input.Content != nil {
		query += ", content = $" + strconv.Itoa(i)
		args = append(args, *input.Content)
		i++
	}
	if input.IsActive != nil {
		query += ", is_active = $" + strconv.Itoa(i)
		args = append(args, *input.IsActive)
		i++
	}

	query += " WHERE banner_id = $" + strconv.Itoa(i)
	args = append(args, bannerID)

	return query, args
}
