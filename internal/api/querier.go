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

func buildGetBannerQuery(params map[string]string) (string, []interface{}) {
	query := "SELECT * FROM banners WHERE True"
	var args []interface{}
	i := 1

	if featureIDStr, ok := params["feature_id"]; ok && featureIDStr != "" {
		featureID, err := strconv.Atoi(featureIDStr)
		if err == nil {
			query += " AND feature_id = $" + strconv.Itoa(i)
			args = append(args, featureID)
			i++
		}
	}

	if tagIDStr, ok := params["tag_id"]; ok && tagIDStr != "" {
		tagID, err := strconv.Atoi(tagIDStr)
		if err == nil {
			query += " AND $" + strconv.Itoa(i) + " = ANY(tag_ids)"
			args = append(args, tagID)
			i++
		}
	}

	if limitStr, ok := params["limit"]; ok && limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil {
			query += " LIMIT $" + strconv.Itoa(i)
			args = append(args, limit)
			i++
		}
	}

	if offsetStr, ok := params["offset"]; ok && offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err == nil {
			query += " OFFSET $" + strconv.Itoa(i)
			args = append(args, offset)
			i++
		}
	}
	return query, args
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
