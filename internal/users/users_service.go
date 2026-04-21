package users

func NewService(store StoreInterface) ServiceInterface {
	return &Service{
		store: store,
	}
}

type Service struct {
	store StoreInterface
}

func (svc *Service) CreateUser(userdata *CreateUserParam) error {
	// input validation

	// hash password

	err := svc.store.Create(userdata)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateUser(userId int64, updatedUserData *UpdateUserParam) error {
	err := svc.store.Update(userId, updatedUserData)
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
