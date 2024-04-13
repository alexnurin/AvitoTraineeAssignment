CREATE TABLE features
(
    feature_id   SERIAL PRIMARY KEY
);

CREATE TABLE tags
(
    tag_id   SERIAL PRIMARY KEY
);

CREATE TABLE banners
(
    banner_id   SERIAL PRIMARY KEY,
    banner_data JSON    NOT NULL,
    feature_id  INTEGER NOT NULL,
    UNIQUE (banner_id, feature_id),
    FOREIGN KEY (feature_id) REFERENCES features (feature_id)
);

CREATE TABLE banner_tags
(
    banner_id INTEGER NOT NULL,
    tag_id    INTEGER NOT NULL,
    PRIMARY KEY (banner_id, tag_id),
    FOREIGN KEY (banner_id) REFERENCES banners (banner_id),
    FOREIGN KEY (tag_id) REFERENCES tags (tag_id)
);

CREATE TABLE users
(
    user_id SERIAL PRIMARY KEY,
    is_vip  BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE user_tags
(
    user_id INTEGER NOT NULL,
    tag_id  INTEGER NOT NULL,
    PRIMARY KEY (user_id, tag_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (tag_id) REFERENCES tags (tag_id)
);

-- CREATE TABLE user_banner_visibility
-- (
--     user_id     INTEGER   NOT NULL,
--     banner_id   INTEGER   NOT NULL,
--     is_current  BOOLEAN   NOT NULL,
--     last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     PRIMARY KEY (user_id, banner_id),
--     FOREIGN KEY (user_id) REFERENCES users (user_id),
--     FOREIGN KEY (banner_id) REFERENCES banners (banner_id)
-- );

-- Запрос для получения актуальных баннеров, включая механизм для VIP-пользователей
-- SELECT b.banner_data
-- FROM banners b
--          JOIN banner_tags bt ON b.banner_id = bt.banner_id
--          JOIN user_tags ut ON bt.tag_id = ut.tag_id
--          JOIN user_banner_visibility ubv ON ut.user_id = ubv.user_id
--          JOIN users u ON ubv.user_id = u.user_id
-- WHERE ubv.user_id = 1
--   AND (ubv.is_current = TRUE OR (u.is_vip = TRUE AND ubv.last_update >= NOW() - INTERVAL '1 day'));
