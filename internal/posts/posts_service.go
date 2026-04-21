package posts

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) GetPosts() (*[]Post, error) {
	posts, err := svc.store.Get()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (svc *Service) GetPostById(postId int64) (*Post, error) {
	posts, err := svc.store.GetById(postId)
	if err != nil {
		return &Post{}, err
	}
	return posts, nil
}

func (svc *Service) CreatePost(postData *CreatePostParam) error {
	err := svc.store.Create(postData)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeletePost(postId int64) error {
	err := svc.store.Delete(postId)
	if err != nil {
		return err
	}
	return nil
}
