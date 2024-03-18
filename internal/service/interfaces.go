package service

import "github.com/akrovv/filmlibrary/internal/domain"

type UserStorage interface {
	Register(user *domain.CRUser) error
	Login(user *domain.CRUser) (*domain.User, error)
}

type SessionStorage interface {
	Create(dto *domain.CreateSession) (string, error)
	Get(dto *domain.GetSession) (*domain.User, error)
}

type ActorStorage interface {
	Create(dto *domain.CreateActor) error
	Update(dto *domain.Actor) error
	Delete(dto *domain.DeleteActor) error
	GetList() ([]domain.ActorWithMovie, error)
}

type MovieStorage interface {
	Create(dto *domain.CreateMovie) error
	Update(dto *domain.Movie) error
	Delete(dto *domain.DeleteMovie) error
	Get(dto *domain.GetMovie) (*domain.MovieWithoudID, error)
	GetOrderedList(dto *domain.GetOrderedMovie) ([]domain.MovieWithoudID, error)
}
