package follows

// Interface
type StoreInterface interface {
	Follow(userId int64, followedUserId int64) error
	Unfollow(userId int64, followedUserId int64) error
	GetFollower(userId int64) (*[]Follow, error)
	GetFollowed(userId int64) (*[]Follow, error)
}

type ServiceInterface interface {
	Follow(userId int64, followedUserId int64) error
	Unfollow(userId int64, followedUserId int64) error
	GetFollower(userId int64) (*[]Follow, error)
	GetFollowed(userId int64) (*[]Follow, error)
}

// struct
type Follow struct {
	ID           int64  `json:"id,omitempty"`
	FollowedId   int64  `json:"followed_id,omitempty"`
	FollowedName string `json:"followed_name,omitempty"`
	FollowerId   int64  `json:"follower_id,omitempty"`
	FollowerName string `json:"follower_name,omitempty"`
}
