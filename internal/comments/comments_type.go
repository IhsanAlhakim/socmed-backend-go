package comments

import "time"

// Interface
type StoreInterface interface {
	CreateComment(postData *CreateCommentParam) error
	Getcomments(postId int64) (*[]Comment, error)
}

type ServiceInterface interface {
	CreateComment(postData *CreateCommentParam) error
	Getcomments(postId int64) (*[]Comment, error)
}

// struct
type Comment struct {
	ID        int64     `json:"id,omitempty"`
	PostId    int64     `json:"post_id,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreateCommentParam struct {
	UserId  int64  `json:"user_id"`
	PostId  int64  `json:"post_id"`
	Content string `json:"content"`
}
