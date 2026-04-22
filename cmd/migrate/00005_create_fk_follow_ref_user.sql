-- +goose Up
-- +goose StatementBegin
ALTER TABLE follows
ADD CONSTRAINT fk_follow_followed_id
FOREIGN KEY (followed_id)
REFERENCES  users (id) ON DELETE CASCADE
;

ALTER TABLE follows
ADD CONSTRAINT fk_follow_follower_id
FOREIGN KEY (follower_id)
REFERENCES  users (id) ON DELETE CASCADE
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE follows
DROP CONSTRAINT fk_follow_followed_id;

ALTER TABLE follows
DROP CONSTRAINT fk_follow_follower_id;
-- +goose StatementEnd
