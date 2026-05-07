-- +goose Up
ALTER TABLE follows
ADD CONSTRAINT follows_followedid_followerid_unique
UNIQUE (followed_id, follower_id)
;

-- +goose Down
ALTER TABLE follows
DROP CONSTRAINT follows_followedid_followerid_unique
;
