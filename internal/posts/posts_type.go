package posts

import "time"

// Interface
type StoreInterface interface {
	CreatePost(userId int64, payload *CreatePostParam) error
	DeletePost(postId int64) error
	GetPosts(userId int64) (*[]Post, error)
	GetFollowedPosts(userId int64) (*[]Post, error)
	GetById(postId int64) (*Post, error)
	GetLikedPosts(userId int64) (*[]Post, error)
	GetPostsByUsername(userId int64, otherUsername string) (*[]Post, error)
}

type ServiceInterface interface {
	CreatePost(userId int64, payload *CreatePostParam) error
	GetPosts(userId int64) (*[]Post, error)
	GetFollowedPosts(userId int64) (*[]Post, error)
	GetPostById(postId int64) (*Post, error)
	DeletePost(postId int64) error
	GetLikedPosts(userId int64) (*[]Post, error)
	GetPostsByUsername(userId int64, otherUsername string) (*[]Post, error)
}

// struct
type Post struct {
	ID        int64     `json:"id,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	Creator   string    `json:"creator,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Liked     bool      `json:"liked"`
}

type CreatePostParam struct {
	Content string `json:"content" validate:"required"`
}
