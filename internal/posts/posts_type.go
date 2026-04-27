package posts

import "time"

// Interface
type StoreInterface interface {
	Create(payload *CreatePostParam) error
	Delete(postId int64) error
	GetPosts() (*[]Post, error)
	GetFollowedPosts(userId int64) (*[]Post, error)
	GetById(postId int64) (*Post, error)
}

type ServiceInterface interface {
	CreatePost(payload *CreatePostParam) error
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
	UserId  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
