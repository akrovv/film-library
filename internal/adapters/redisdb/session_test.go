package redisdb

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/pkg/hasher"
	"github.com/go-redis/redismock/v9"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	client, mock := redismock.NewClientMock()
	defer client.Close()

	badHasher := hasher.NewBadHasher()
	hasher := hasher.NewHasher([]byte("session"))
	storage := NewSessionStorage(ctx, client, hasher)

	hashedNickname, err := hasher.GetHash("user")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	createSession := domain.CreateSession{
		Username: "user",
		IsAdmin:  false,
	}

	data, err := json.Marshal(createSession)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// OK
	mock.ExpectSet(hashedNickname, data, time.Hour*8).SetVal(hashedNickname)
	sessionID, err := storage.Create(&createSession)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !reflect.DeepEqual(sessionID, hashedNickname) {
		t.Errorf("expected: %v, got: %v", hashedNickname, sessionID)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Redis returned error
	mock.ExpectSet(hashedNickname, data, time.Hour*8).SetErr(domain.ErrTest)
	sessionID, err = storage.Create(&createSession)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if sessionID != "" {
		t.Errorf("expected empty string, got: %s", sessionID)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// GetHash returned error
	storage = NewSessionStorage(ctx, client, badHasher)
	_, err = storage.Create(&createSession)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	client, mock := redismock.NewClientMock()
	defer client.Close()

	hasher := hasher.NewHasher([]byte("session"))
	storage := NewSessionStorage(ctx, client, hasher)

	hashedNickname, err := hasher.GetHash("user")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	expectedUser := &domain.User{
		Username: "user",
		IsAdmin:  false,
	}

	getSession := domain.GetSession{
		Username: hashedNickname,
	}

	// OK
	mock.ExpectGet(hashedNickname).SetVal(`{"username": "user", "is_admin": "false"}`)
	user, err := storage.Get(&getSession)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("expected: %v, got: %v", expectedUser, user)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Redis returned error
	mock.ExpectGet(hashedNickname).SetErr(domain.ErrTest)
	user, err = storage.Get(&getSession)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if user != nil {
		t.Errorf("expected nil, got: %v", user)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Redis returned incorrect json
	mock.ExpectGet(hashedNickname).SetVal(`incorrect json`)
	user, err = storage.Get(&getSession)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if user != nil {
		t.Errorf("expected nil, got: %v", user)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
