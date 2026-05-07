package plikes

// Interface
type StoreInterface interface {
	LikePost(postId int64, userId int64) error
	UnlikePost(postId int64, userId int64) error
	GetPostLiker(postId int64) (*[]PostLike, error)
	GetPostLikesCount(postId int64) (*PostLikesCount, error)
}

type ServiceInterface interface {
	LikePost(postId int64, userId int64) error
	UnlikePost(postId int64, userId int64) error
	GetPostLiker(postId int64) (*[]PostLike, error)
	GetPostLikesCount(postId int64) (*PostLikesCount, error)
}

// struct
type PostLike struct {
	ID       int64  `json:"id,omitempty"`
	PostId   int64  `json:"post_id,omitempty"`
	UserId   int64  `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type PostLikesCount struct {
	LikesCount int32 `json:"likes_count"`
}
