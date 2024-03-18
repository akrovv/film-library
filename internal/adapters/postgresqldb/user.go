package postgresqldb

import (
	"database/sql"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/pkg/hasher"
)

type userStorage struct {
	db     *sql.DB
	hasher hasher.Hasher
}

func NewUserStorage(db *sql.DB, hasher hasher.Hasher) *userStorage {
	return &userStorage{
		db:     db,
		hasher: hasher,
	}
}

func (s *userStorage) Register(user *domain.CRUser) error {
	hashedPassword, err := s.hasher.GetHash(user.Password)
	if err != nil {
		return err
	}

	_, err = s.db.Exec("INSERT INTO Users (username, password) VALUES ($1, $2)", user.Username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (s *userStorage) Login(user *domain.CRUser) (*domain.User, error) {
	hashedPassword, err := s.hasher.GetHash(user.Password)
	if err != nil {
		return nil, err
	}

	curUser := domain.User{}
	if err = s.db.QueryRow("SELECT username, is_admin FROM Users WHERE username=$1 AND password=$2", user.Username,
		hashedPassword).Scan(&curUser.Username, &curUser.IsAdmin); err != nil {
		return nil, err
	}

	return &curUser, nil
}
