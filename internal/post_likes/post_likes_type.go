package plikes

// Interface
type StoreInterface interface {
	LikePost(payload *PostLikeParam) error
	UnlikePost(payload *PostLikeParam) error
	GetPostLiker(postId int64) (*[]PostLike, error)
}

type ServiceInterface interface {
	LikePost(payload *PostLikeParam) error
	UnlikePost(payload *PostLikeParam) error
	GetPostLiker(postId int64) (*[]PostLike, error)
}

// struct
type PostLike struct {
	ID       int64  `json:"id,omitempty"`
	PostId   int64  `json:"post_id,omitempty"`
	UserId   int64  `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type PostLikeParam struct {
	PostId int64 `json:"post_id,omitempty"`
	UserId int64 `json:"user_id,omitempty"`
}
