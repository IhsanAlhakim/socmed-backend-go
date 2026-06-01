-- +goose Up
ALTER TABLE posts
DROP COLUMN title;

-- +goose Down
ALTER TABLE posts
ADD COLUMN title TEXT NOT NULL;
