package service

import "github.com/akrovv/filmlibrary/internal/domain"

type userService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *userService {
	return &userService{
		storage: storage,
	}
}

func (s *userService) Register(user *domain.CRUser) error {
	return s.storage.Register(user)
}

func (s *userService) Login(user *domain.CRUser) (*domain.User, error) {
	return s.storage.Login(user)
}
