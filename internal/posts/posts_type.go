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
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatePostParam struct {
	UserId  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
