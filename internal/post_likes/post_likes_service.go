package plikes

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) LikePost(payload *PostLikeParam) error {
	err := svc.store.LikePost(payload)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UnlikePost(payload *PostLikeParam) error {
	err := svc.store.UnlikePost(payload)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) GetPostLiker(postId int64) (*[]PostLike, error) {
	postLiker, err := svc.store.GetPostLiker(postId)
	if err != nil {
		return nil, err
	}
	return postLiker, nil
}
