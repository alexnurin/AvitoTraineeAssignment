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

