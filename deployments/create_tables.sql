CREATE TABLE IF NOT EXISTS banners
(
    banner_id  SERIAL PRIMARY KEY,
    content    JSONB NOT NULL,
    feature_id INT   NOT NULL,
    tag_ids    INT[]     DEFAULT '{}',
    is_active  BOOLEAN   DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- CREATE INDEX idx_feature_id ON banners (feature_id);
-- CREATE INDEX idx_is_active ON banners (is_active);
-- CREATE INDEX idx_tag_ids ON banners USING GIN (tag_ids);
-- CREATE INDEX idx_feature_id_is_active ON banners (feature_id, is_active);


