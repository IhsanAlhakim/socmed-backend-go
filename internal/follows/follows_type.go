package follows

// Interface
type StoreInterface interface {
	Create(payload *FollowParam) error
	Delete(payload *FollowParam) error
	GetFollower(followedId int64) (*[]Follow, error)
	GetFollowed(followerId int64) (*[]Follow, error)
}

type ServiceInterface interface {
	Follow(payload *FollowParam) error
	Unfollow(payload *FollowParam) error
	GetFollower(followedId int64) (*[]Follow, error)
	GetFollowed(followerId int64) (*[]Follow, error)
}

// struct
type Follow struct {
	ID           int64  `json:"id,omitempty"`
	FollowedId   int64  `json:"followed_id,omitempty"`
	FollowedName string `json:"followed_name,omitempty"`
	FollowerId   int64  `json:"follower_id,omitempty"`
	FollowerName string `json:"follower_name,omitempty"`
}

type FollowParam struct {
	FollowedId int64 `json:"followed_id"`
	FollowerId int64 `json:"follower_id"`
}
