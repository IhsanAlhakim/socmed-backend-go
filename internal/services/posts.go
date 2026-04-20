package services

import (
	"github.com/IhsanAlhakim/socmed-backend-go/internal/store"
)

type PostServiceInterface interface {
	CreatePost(postData store.Post) error
}

type PostService struct {
	storage store.Storage
}

func NewPostService(storage store.Storage) *PostService {
	return &PostService{
		storage: storage,
	}
}

func (psvc *PostService) CreatePost(postData store.Post) error {
	err := psvc.storage.Posts.Create(&postData)
	if err != nil {
		return err
	}
	return nil
}

func (psvc *PostService) DeletePost(postId int64) error {
	err := psvc.storage.Posts.Delete(postId)
	if err != nil {
		return err
	}
	return nil
}
