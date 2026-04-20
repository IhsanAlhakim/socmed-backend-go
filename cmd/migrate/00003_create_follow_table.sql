-- +goose Up
CREATE TABLE IF NOT EXISTS follows (
    id BIGSERIAL PRIMARY KEY,
    followed_id BIGINT NOT NULL,
    follower_id BIGINT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS follows;
