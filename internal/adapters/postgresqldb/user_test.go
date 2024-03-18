package postgresqldb

import (
	"reflect"
	"testing"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/pkg/hasher"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUserCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	user := &domain.CRUser{
		Username: "user",
		Password: "user",
	}

	badHasher := hasher.NewBadHasher()
	hasher := hasher.NewHasher([]byte("secret"))
	hashedPassword, err := hasher.GetHash(user.Password)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	storage := NewUserStorage(db, hasher)

	// OK
	mock.ExpectExec(`INSERT INTO Users \(username, password\) VALUES \(\$1, \$2\)`).
		WithArgs(user.Username, hashedPassword).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.Register(user)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres returned error
	mock.ExpectExec(`INSERT INTO Users \(username, password\) VALUES \(\$1, \$2\)`).
		WithArgs(user.Username, hashedPassword).
		WillReturnError(domain.ErrTest)

	err = storage.Register(user)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// GetHash returned error
	storage = NewUserStorage(db, badHasher)

	err = storage.Register(user)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	expectedUser := &domain.User{
		Username: "user",
	}

	userTest := &domain.CRUser{
		Username: "user",
		Password: "user",
	}

	badHasher := hasher.NewBadHasher()
	hasher := hasher.NewHasher([]byte("secret"))
	hashedPassword, err := hasher.GetHash("user")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	storage := NewUserStorage(db, hasher)

	// OK
	mock.ExpectQuery("SELECT username, is_admin FROM Users WHERE username=\\$1 AND password=\\$2").
		WithArgs(expectedUser.Username, hashedPassword).
		WillReturnRows(sqlmock.NewRows([]string{"username", "is_admin"}).
			AddRow(expectedUser.Username, false))

	user, err := storage.Login(userTest)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("expected: %v, got: %v", expectedUser, user)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres returned error
	mock.ExpectQuery("SELECT username, is_admin FROM Users WHERE username=\\$1 AND password=\\$2").
		WithArgs(expectedUser.Username, hashedPassword).
		WillReturnError(domain.ErrTest)

	user, err = storage.Login(userTest)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if user != nil {
		t.Errorf("expected nil, got: %v", user)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// GetHash returned error
	storage = NewUserStorage(db, badHasher)

	user, err = storage.Login(userTest)
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
