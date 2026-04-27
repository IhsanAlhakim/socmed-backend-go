package comments

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) CreateComment(postId int64, payload *CreateCommentParam) error {
	err := svc.store.CreateComment(postId, payload)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Getcomments(postId int64) (*[]Comment, error) {
	comments, err := svc.store.Getcomments(postId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (svc *Service) DeleteComment(commentId int64) error {
	err := svc.store.DeleteComment(commentId)
	if err != nil {
		return err
	}
	return nil
}
