package service

import "github.com/akrovv/filmlibrary/internal/domain"

type sessionService struct {
	storage SessionStorage
}

func NewSessionService(storage SessionStorage) *sessionService {
	return &sessionService{
		storage: storage,
	}
}

func (s *sessionService) Create(dto *domain.CreateSession) (string, error) {
	return s.storage.Create(dto)
}

func (s *sessionService) Get(dto *domain.GetSession) (*domain.User, error) {
	return s.storage.Get(dto)
}
