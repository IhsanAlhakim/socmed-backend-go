package posts

import (
	"time"
)

type Post struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Interface
type StoreInterface interface {
	Create(postData *Post) error
	Delete(postId int64) error
	Get() (*[]Post, error)
	GetById(postId int64) (*Post, error)
}

type ServiceInterface interface {
	CreatePost(postData *Post) error
	GetPosts() (*[]Post, error)
	GetPostById(postId int64) (*Post, error)
	DeletePost(postId int64) error
}
