package service

import "github.com/akrovv/filmlibrary/internal/domain"

type actorService struct {
	storage ActorStorage
}

func NewActorService(storage ActorStorage) *actorService {
	return &actorService{
		storage: storage,
	}
}

func (s *actorService) Create(dto *domain.CreateActor) error {
	return s.storage.Create(dto)
}

func (s *actorService) Update(dto *domain.Actor) error {
	return s.storage.Update(dto)
}

func (s *actorService) Delete(dto *domain.DeleteActor) error {
	return s.storage.Delete(dto)
}

func (s *actorService) GetList() ([]domain.ActorWithMovie, error) {
	return s.storage.GetList()
}
