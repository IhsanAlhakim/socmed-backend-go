package follows

type Follow struct {
	ID         int64 `json:"id"`
	FollowedId int64 `json:"followed_id"`
	FollowerId int64 `json:"follower_id"`
}

type StoreInterface interface {
	Create(followData *Follow) error
	Delete(followData *Follow) error
	Get(followedId int64) (*[]Follow, error)
}

type ServiceInterface interface {
	Follow(followData *Follow) error
	Unfollow(followData *Follow) error
	GetFollower(followedId int64) (*[]Follow, error)
}
