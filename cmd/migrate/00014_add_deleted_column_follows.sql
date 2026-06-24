-- +goose Up
ALTER TABLE follows
ADD COLUMN deleted BOOLEAN NOT NULL DEFAULT FALSE;


-- +goose Down
ALTER TABLE follows
DROP COLUMN deleted;

