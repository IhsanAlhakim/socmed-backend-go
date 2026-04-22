package follows

// Interface
type StoreInterface interface {
	Create(followData *FollowDataparam) error
	Delete(followData *FollowDataparam) error
	GetFollower(followedId int64) (*[]Follow, error)
	GetFollowed(followerId int64) (*[]Follow, error)
}

type ServiceInterface interface {
	Follow(followData *FollowDataparam) error
	Unfollow(followData *FollowDataparam) error
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

type FollowDataparam struct {
	FollowedId int64 `json:"followed_id"`
	FollowerId int64 `json:"follower_id"`
}
