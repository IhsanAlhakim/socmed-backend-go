package posts

import "time"

// Interface
type StoreInterface interface {
	Create(postData *CreatePostParam) error
	Delete(postId int64) error
	Get() (*[]Post, error)
	GetById(postId int64) (*Post, error)
}

type ServiceInterface interface {
	CreatePost(postData *CreatePostParam) error
	GetPosts() (*[]Post, error)
	GetPostById(postId int64) (*Post, error)
	DeletePost(postId int64) error
}

// struct
type Post struct {
	ID        int64     `json:"id,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreatePostParam struct {
	UserId  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
