-- +goose Up
ALTER TABLE post_likes
ADD CONSTRAINT post_likes_postId_userId_unique
UNIQUE (post_id, user_id)
;

-- +goose Down
ALTER TABLE post_likes
DROP CONSTRAINT post_likes_postId_userId_unique
;
