package restapi

import (
	"github.com/akrovv/filmlibrary/internal/domain"
)

type UserService interface {
	Register(user *domain.CRUser) error
	Login(user *domain.CRUser) (*domain.User, error)
}

type SessionService interface {
	Create(dto *domain.CreateSession) (string, error)
	Get(dto *domain.GetSession) (*domain.User, error)
}

type ActorService interface {
	Create(dto *domain.CreateActor) error
	Update(dto *domain.Actor) error
	Delete(dto *domain.DeleteActor) error
	GetList() ([]domain.ActorWithMovie, error)
}

type MovieService interface {
	Create(dto *domain.CreateMovie) error
	Update(dto *domain.Movie) error
	Delete(dto *domain.DeleteMovie) error
	Get(dto *domain.GetMovie) (*domain.MovieWithoudID, error)
	GetOrderedList(dto *domain.GetOrderedMovie) ([]domain.MovieWithoudID, error)
}
