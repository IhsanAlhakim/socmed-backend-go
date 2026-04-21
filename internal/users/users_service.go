package users

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) CreateUser(userdata *User) error {
	// input validation

	// hash password

	err := svc.store.Create(userdata)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateUser(userId int64, userdata *User) error {
	err := svc.store.Update(userId, userdata)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeleteUser(userId int64) error {

	err := svc.store.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
