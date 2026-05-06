package follows

import "github.com/IhsanAlhakim/socmed-backend-go/internal/validation"

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) GetFollower(userId int64) (*[]Follow, error) {
	follower, err := svc.store.GetFollower(userId)
	if err != nil {
		return nil, err
	}
	return follower, nil
}

func (svc *Service) GetFollowed(userId int64) (*[]Follow, error) {
	followed, err := svc.store.GetFollowed(userId)
	if err != nil {
		return nil, err
	}
	return followed, nil
}

func (svc *Service) Follow(userId int64, payload *FollowParam) error {
	if err := validation.Validate.Struct(payload); err != nil {
		return err
	}
	err := svc.store.Follow(userId, payload)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Unfollow(userId int64, payload *FollowParam) error {
	if err := validation.Validate.Struct(payload); err != nil {
		return err
	}
	err := svc.store.Unfollow(userId, payload)
	if err != nil {
		return err
	}
	return nil
}
