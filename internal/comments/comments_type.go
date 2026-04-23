package comments

import "time"

// Interface
type StoreInterface interface {
	CreateComment(postData *CreateCommentParam) error
}

type ServiceInterface interface {
	CreateComment(postData *CreateCommentParam) error
}

// struct
type Comment struct {
	ID        int64     `json:"id,omitempty"`
	PostId    int64     `json:"post_id,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreateCommentParam struct {
	UserId  int64  `json:"user_id"`
	PostId  int64  `json:"post_id"`
	Content string `json:"content"`
}
