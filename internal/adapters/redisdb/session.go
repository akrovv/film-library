package redisdb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/pkg/hasher"
	"github.com/redis/go-redis/v9"
)

type sessionStorage struct {
	ctx    context.Context
	db     *redis.Client
	hasher hasher.Hasher
}

func NewSessionStorage(ctx context.Context, db *redis.Client, hasher hasher.Hasher) *sessionStorage {
	return &sessionStorage{
		ctx:    ctx,
		db:     db,
		hasher: hasher,
	}
}

func (s *sessionStorage) Create(dto *domain.CreateSession) (string, error) {
	hashedUsername, err := s.hasher.GetHash(dto.Username)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}

	status := s.db.Set(s.ctx, hashedUsername, data, time.Hour*8)

	if status.Err() != nil {
		return "", status.Err()
	}

	return hashedUsername, nil
}

func (s *sessionStorage) Get(dto *domain.GetSession) (*domain.User, error) {
	user := domain.User{}

	value, err := s.db.Get(s.ctx, dto.Username).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
