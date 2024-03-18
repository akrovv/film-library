package postgresqldb

import (
	"reflect"
	"testing"
	"time"

	"github.com/akrovv/filmlibrary/internal/domain"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestActorCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	storage := NewActorStorage(db)
	actor := domain.CreateActor{
		Name:      "user",
		Gender:    "man",
		DateBirth: time.Now(),
	}

	// OK
	mock.ExpectExec(`INSERT INTO Actors \(actor_name, gender, date_of_birth\) VALUES \(\$1, \$2, \$3\)`).
		WithArgs(actor.Name, actor.Gender, actor.DateBirth).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.Create(&actor)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres returned error
	mock.ExpectExec(`INSERT INTO Actors \(actor_name, gender, date_of_birth\) VALUES \(\$1, \$2, \$3\)`).
		WithArgs(actor.Name, actor.Gender, actor.DateBirth).
		WillReturnError(domain.ErrTest)

	err = storage.Create(&actor)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestActorUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	storage := NewActorStorage(db)
	actor := domain.Actor{
		ID:        1,
		Name:      "user",
		Gender:    "man",
		DateBirth: time.Now(),
	}

	// OK
	mock.ExpectExec(`UPDATE Actors SET actor_name=\$1, gender=\$2, date_of_birth=\$3 WHERE actor_id=\$4`).
		WithArgs(actor.Name, actor.Gender, actor.DateBirth, actor.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.Update(&actor)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres returned error
	mock.ExpectExec(`UPDATE Actors SET actor_name=\$1, gender=\$2, date_of_birth=\$3 WHERE actor_id=\$4`).
		WithArgs(actor.Name, actor.Gender, actor.DateBirth, actor.ID).
		WillReturnError(domain.ErrTest)

	err = storage.Update(&actor)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestActorDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	storage := NewActorStorage(db)
	actor := domain.DeleteActor{
		ID: 1,
	}

	// OK
	mock.ExpectExec(`DELETE FROM Actors WHERE actor_id=\$1`).
		WithArgs(actor.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.Delete(&actor)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres returned error
	mock.ExpectExec(`DELETE FROM Actors WHERE actor_id=\$1`).
		WithArgs(actor.ID).
		WillReturnError(domain.ErrTest)

	err = storage.Delete(&actor)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestActorGetList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()
	storage := NewActorStorage(db)

	testTime := time.Now()
	expectedActors := []domain.ActorWithMovie{
		{CreateActor: domain.CreateActor{
			Name: "Actor 1", Gender: "Male", DateBirth: testTime,
		},
			Title: "Movie 1"},
		{
			CreateActor: domain.CreateActor{
				Name: "Actor 2", Gender: "Female", DateBirth: testTime,
			},
			Title: "Movie 2",
		},
	}

	rows := sqlmock.NewRows([]string{"actor_name", "gender", "date_of_birth", "movie_title"}).
		AddRow("Actor 1", "Male", testTime, "Movie 1").
		AddRow("Actor 2", "Female", testTime, "Movie 2")

	rowsWithError := sqlmock.NewRows([]string{"actor_name", "gender"}).
		AddRow("Actor 1", "Male")

	// OK
	mock.ExpectQuery(`SELECT actor_name, gender, date_of_birth, movie_title FROM Actors JOIN MovieActors USING\(actor_id\) JOIN Movies USING\(movie_id\)`).
		WillReturnRows(rows)

	actors, err := storage.GetList()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !reflect.DeepEqual(actors, expectedActors) {
		t.Errorf("expected: %v, got: %v", expectedActors, actors)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres returned error
	mock.ExpectQuery(`SELECT actor_name, gender, date_of_birth, movie_title FROM Actors JOIN MovieActors USING\(actor_id\) JOIN Movies USING\(movie_id\)`).
		WillReturnError(domain.ErrTest)

	actors, err = storage.GetList()

	if err == nil {
		t.Error("expected error, got nil")
	}

	if actors != nil {
		t.Errorf("expected nil, got: %v", actors)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Rows scan error
	mock.ExpectQuery(`SELECT actor_name, gender, date_of_birth, movie_title FROM Actors JOIN MovieActors USING\(actor_id\) JOIN Movies USING\(movie_id\)`).
		WillReturnRows(rowsWithError)

	actors, err = storage.GetList()

	if err == nil {
		t.Error("expected error, got nil")
	}

	if actors != nil {
		t.Errorf("expected nil, got: %v", actors)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
