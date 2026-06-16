package posts

import "github.com/IhsanAlhakim/socmed-backend-go/internal/validation"

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) GetPosts(userId int64) (*[]Post, error) {
	posts, err := svc.store.GetPosts(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (svc *Service) GetPostsByUsername(userId int64, otherUsername string) (*[]Post, error) {
	posts, err := svc.store.GetPostsByUsername(userId, otherUsername)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (svc *Service) GetFollowedPosts(userId int64) (*[]Post, error) {
	posts, err := svc.store.GetFollowedPosts(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (svc *Service) GetLikedPosts(userId int64) (*[]Post, error) {
	posts, err := svc.store.GetLikedPosts(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (svc *Service) GetPostById(postId int64) (*Post, error) {
	posts, err := svc.store.GetById(postId)
	if err != nil {
		return &Post{}, err
	}
	return posts, nil
}

func (svc *Service) CreatePost(userId int64, payload *CreatePostParam) error {
	if err := validation.Validate.Struct(payload); err != nil {
		return validation.NewError(err)
	}

	err := svc.store.CreatePost(userId, payload)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeletePost(postId int64) error {
	err := svc.store.DeletePost(postId)
	if err != nil {
		return err
	}
	return nil
}
