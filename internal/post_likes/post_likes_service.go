package plikes

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) LikePost(postId int64, payload *PostLikeParam) error {
	err := svc.store.LikePost(postId, payload)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UnlikePost(postId int64, payload *PostLikeParam) error {
	err := svc.store.UnlikePost(postId, payload)
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
