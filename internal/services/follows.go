package services

import (
	"github.com/IhsanAlhakim/socmed-backend-go/internal/store"
)

type FollowServiceInterface interface {
	Follow(followData *store.Follow) error
}

type FollowService struct {
	storage store.Storage
}

func NewFollowService(storage store.Storage) *FollowService {
	return &FollowService{
		storage: storage,
	}
}

func (fsvc *FollowService) Follow(followData *store.Follow) error {
	err := fsvc.storage.Follows.Create(followData)
	if err != nil {
		return err
	}
	return nil
}
