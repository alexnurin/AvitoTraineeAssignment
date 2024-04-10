-- Создание таблицы для хранения информации о баннерах для пользователей
CREATE TABLE user_banners
(
    id                SERIAL PRIMARY KEY,
    tag_id            INTEGER      NOT NULL,
    feature_id        INTEGER      NOT NULL,
    use_last_revision BOOLEAN      NOT NULL DEFAULT false,
    token             VARCHAR(255) NOT NULL,
    created_at        TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_banners_tag_id FOREIGN KEY (tag_id) REFERENCES tags (id),
    CONSTRAINT fk_user_banners_feature_id FOREIGN KEY (feature_id) REFERENCES features (id)
);

-- Создание таблицы для хранения информации о баннерах
CREATE TABLE banners
(
    id         SERIAL PRIMARY KEY,
    feature_id INTEGER NOT NULL,
    content    JSONB   NOT NULL,
    is_active  BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_banners_feature_id FOREIGN KEY (feature_id) REFERENCES features (id)
);

-- Таблица тэгов
CREATE TABLE tags
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Таблица фич
CREATE TABLE features
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Таблица для связи баннеров и тегов
CREATE TABLE banner_tags
(
    banner_id INTEGER NOT NULL,
    tag_id    INTEGER NOT NULL,
    PRIMARY KEY (banner_id, tag_id),
    CONSTRAINT fk_banner_tags_banner_id FOREIGN KEY (banner_id) REFERENCES banners (id),
    CONSTRAINT fk_banner_tags_tag_id FOREIGN KEY (tag_id) REFERENCES tags (id)
);
