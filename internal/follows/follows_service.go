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

func (svc *Service) GetFollowed(followerId int64) (*[]Follow, error) {
	followed, err := svc.store.GetFollowed(followerId)
	if err != nil {
		return nil, err
	}
	return followed, nil
}

func (svc *Service) Follow(payload *FollowParam) error {
	err := svc.store.Create(payload)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Unfollow(payload *FollowParam) error {
	err := svc.store.Delete(payload)
	if err != nil {
		return err
	}
	return nil
}
