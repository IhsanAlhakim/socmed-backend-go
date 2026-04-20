package services

import (
	"github.com/IhsanAlhakim/socmed-backend-go/internal/store"
)

type PostServiceInterface interface {
	CreatePost(postData store.Post) error
	GetPosts() ([]store.Post, error)
	GetPostById(postId int64) (store.Post, error)
}

type PostService struct {
	storage store.Storage
}

func NewPostService(storage store.Storage) *PostService {
	return &PostService{
		storage: storage,
	}
}

func (psvc *PostService) GetPosts() ([]store.Post, error) {
	posts, err := psvc.storage.Posts.Get()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (psvc *PostService) GetPostById(postId int64) (store.Post, error) {
	posts, err := psvc.storage.Posts.GetById(postId)
	if err != nil {
		return store.Post{}, err
	}
	return posts, nil
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
