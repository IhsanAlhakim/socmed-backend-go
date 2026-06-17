-- +goose Up
ALTER TABLE post_likes
ADD COLUMN deleted BOOLEAN NOT NULL DEFAULT FALSE;


-- +goose Down
ALTER TABLE post_likes
DROP COLUMN deleted;

