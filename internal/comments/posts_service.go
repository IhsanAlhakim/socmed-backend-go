package comments

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) CreateComment(commentData *CreateCommentParam) error {
	err := svc.store.CreateComment(commentData)
	if err != nil {
		return err
	}
	return nil
}
