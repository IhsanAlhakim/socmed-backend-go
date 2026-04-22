package follows

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) GetFollower(followedId int64) (*[]Follow, error) {
	follower, err := svc.store.GetFollower(followedId)
	if err != nil {
		return nil, err
	}
	return follower, nil
}

func (svc *Service) Follow(followData *FollowDataparam) error {
	err := svc.store.Create(followData)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Unfollow(followData *FollowDataparam) error {
	err := svc.store.Delete(followData)
	if err != nil {
		return err
	}
	return nil
}
