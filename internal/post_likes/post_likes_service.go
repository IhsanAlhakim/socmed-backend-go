package plikes

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) LikePost(postLikeData *PostLikeParam) error {
	err := svc.store.LikePost(postLikeData)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UnlikePost(postLikeData *PostLikeParam) error {
	err := svc.store.UnlikePost(postLikeData)
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
