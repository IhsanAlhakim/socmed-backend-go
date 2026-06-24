package follows

import "errors"

var (
	ErrFollowedNotFound       = errors.New("followed user not found")
	ErrUserAlreadyFollowed    = errors.New("this user is already followed")
	ErrFollowerSameAsFollowed = errors.New("follower_id is the same as followed_id")
)
