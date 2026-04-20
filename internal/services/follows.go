package services

import (
	"github.com/IhsanAlhakim/socmed-backend-go/internal/store"
)

type FollowServiceInterface interface {
	Follow(followData *store.Follow) error
	Unfollow(followData *store.Follow) error
	GetFollower(followedId int64) ([]store.Follow, error)
}

type FollowService struct {
	storage store.Storage
}

func NewFollowService(storage store.Storage) *FollowService {
	return &FollowService{
		storage: storage,
	}
}

func (fsvc *FollowService) GetFollower(followedId int64) ([]store.Follow, error) {
	follower, err := fsvc.storage.Follows.Get(followedId)
	if err != nil {
		return nil, err
	}
	return follower, nil
}

func (fsvc *FollowService) Follow(followData *store.Follow) error {
	err := fsvc.storage.Follows.Create(followData)
	if err != nil {
		return err
	}
	return nil
}

func (fsvc *FollowService) Unfollow(followData *store.Follow) error {
	err := fsvc.storage.Follows.Delete(followData)
	if err != nil {
		return err
	}
	return nil
}
