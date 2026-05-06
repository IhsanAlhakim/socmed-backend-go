package posts

import "time"

// Interface
type StoreInterface interface {
	CreatePost(userId int64, payload *CreatePostParam) error
	DeletePost(postId int64) error
	GetPosts() (*[]Post, error)
	GetFollowedPosts(userId int64) (*[]Post, error)
	GetById(postId int64) (*Post, error)
}

type ServiceInterface interface {
	CreatePost(userId int64, payload *CreatePostParam) error
	GetPosts() (*[]Post, error)
	GetFollowedPosts(userId int64) (*[]Post, error)
	GetPostById(postId int64) (*Post, error)
	DeletePost(postId int64) error
}

// struct
type Post struct {
	ID        int64     `json:"id,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	Creator   string    `json:"creator,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreatePostParam struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
