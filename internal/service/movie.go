package service

import "github.com/akrovv/filmlibrary/internal/domain"

type movieService struct {
	storage MovieStorage
}

func NewMovieService(storage MovieStorage) *movieService {
	return &movieService{
		storage: storage,
	}
}

func (s *movieService) Create(dto *domain.CreateMovie) error {
	return s.storage.Create(dto)
}

func (s *movieService) Update(dto *domain.Movie) error {
	return s.storage.Update(dto)
}

func (s *movieService) Delete(dto *domain.DeleteMovie) error {
	return s.storage.Delete(dto)
}

func (s *movieService) Get(dto *domain.GetMovie) (*domain.MovieWithoudID, error) {
	return s.storage.Get(dto)
}

func (s *movieService) GetOrderedList(dto *domain.GetOrderedMovie) ([]domain.MovieWithoudID, error) {
	return s.storage.GetOrderedList(dto)
}
