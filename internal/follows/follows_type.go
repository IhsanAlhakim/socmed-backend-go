package follows

type Follow struct {
	ID         int64 `json:"id"`
	FollowedId int64 `json:"followed_id"`
	FollowerId int64 `json:"follower_id"`
}

type StoreInterface interface {
	Create(followData *FollowDataparam) error
	Delete(followData *FollowDataparam) error
	Get(followedId int64) (*[]Follow, error)
}

type ServiceInterface interface {
	Follow(followData *FollowDataparam) error
	Unfollow(followData *FollowDataparam) error
	GetFollower(followedId int64) (*[]Follow, error)
}

type FollowDataparam struct {
	FollowedId int64 `json:"followed_id"`
	FollowerId int64 `json:"follower_id"`
}
