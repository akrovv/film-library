package postgresqldb

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/akrovv/filmlibrary/internal/domain"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestMovieCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	storage := NewMovieStorage(db)

	dto := &domain.CreateMovie{
		Title:       "Test Movie",
		Description: "Test Description",
		ReleaseDate: time.Now(),
		Rating:      5,
		Actors:      []int64{1, 2},
	}

	// OK
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO Movies \(movie_title, description, release_date, rating\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING movie_i`).
		WithArgs(dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating).
		WillReturnRows(sqlmock.NewRows([]string{"movie_id"}).AddRow(1))
	mock.ExpectExec(`INSERT INTO MovieActors \(movie_id, actor_id\) VALUES \(\$1, \$2\), \(\$3, \$4\)`).
		WithArgs(1, 1, 1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = storage.Create(dto)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres Begin() returned error
	mock.ExpectBegin().WillReturnError(domain.ErrTest)
	err = storage.Create(dto)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres first CMD returned error
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO Movies \(movie_title, description, release_date, rating\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING movie_i`).
		WithArgs(dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating).
		WillReturnError(domain.ErrTest)
	mock.ExpectRollback()

	err = storage.Create(dto)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres second CMD returned error
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO Movies \(movie_title, description, release_date, rating\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING movie_i`).
		WithArgs(dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating).
		WillReturnRows(sqlmock.NewRows([]string{"movie_id"}).AddRow(1))
	mock.ExpectExec(`INSERT INTO MovieActors \(movie_id, actor_id\) VALUES \(\$1, \$2\), \(\$3, \$4\)`).
		WithArgs(1, 1, 1, 2).
		WillReturnError(domain.ErrTest)
	mock.ExpectRollback()

	err = storage.Create(dto)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMovieUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	storage := NewMovieStorage(db)

	dto := &domain.Movie{
		ID:          1,
		Title:       "Test Movie",
		Description: "Test Description",
		ReleaseDate: time.Now(),
		Rating:      5,
		Actors:      []int64{1, 2},
	}

	// OK
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE Movies SET movie_title = \$1, description = \$2, release_date = \$3, rating = \$4 WHERE movie_id = \$5`).
		WithArgs(dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating, dto.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`DELETE FROM MovieActors WHERE movie_id = \$1`).
		WithArgs(dto.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO MovieActors \(movie_id, actor_id\) VALUES \(\$1, \$2\), \(\$3, \$4\)`).
		WithArgs(1, 1, 1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = storage.Update(dto)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres Begin() returned error
	mock.ExpectBegin().WillReturnError(domain.ErrTest)
	err = storage.Update(dto)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres first CMD returned error
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE Movies SET movie_title = \$1, description = \$2, release_date = \$3, rating = \$4 WHERE movie_id = \$5`).
		WithArgs(dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating, dto.ID).
		WillReturnError(domain.ErrTest)
	mock.ExpectRollback()

	err = storage.Update(dto)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres second CMD returned error
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE Movies SET movie_title = \$1, description = \$2, release_date = \$3, rating = \$4 WHERE movie_id = \$5`).
		WithArgs(dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating, dto.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`DELETE FROM MovieActors WHERE movie_id = \$1`).
		WithArgs(dto.ID).
		WillReturnError(domain.ErrTest)
	mock.ExpectRollback()

	err = storage.Update(dto)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres third CMD returned error
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE Movies SET movie_title = \$1, description = \$2, release_date = \$3, rating = \$4 WHERE movie_id = \$5`).
		WithArgs(dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating, dto.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`DELETE FROM MovieActors WHERE movie_id = \$1`).
		WithArgs(dto.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO MovieActors \(movie_id, actor_id\) VALUES \(\$1, \$2\), \(\$3, \$4\)`).
		WithArgs(1, 1, 1, 2).
		WillReturnError(domain.ErrTest)
	mock.ExpectRollback()

	err = storage.Update(dto)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMovieDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
		return
	}
	defer db.Close()

	storage := NewMovieStorage(db)
	deleteMovie := domain.DeleteMovie{
		ID: 1,
	}

	// OK
	mock.ExpectExec(`DELETE FROM Movies WHERE movie_id = \$1`).
		WithArgs(deleteMovie.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = storage.Delete(&deleteMovie)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// Postgres returned error
	mock.ExpectExec(`DELETE FROM Movies WHERE movie_id = \$1`).
		WithArgs(deleteMovie.ID).
		WillReturnError(domain.ErrTest)

	err = storage.Delete(&deleteMovie)

	if err == nil {
		t.Error("expected error, got nil")
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestMovieGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	storage := NewMovieStorage(db)

	dto1 := &domain.GetMovie{
		Title:      "Test Movie",
		SearchType: "title",
	}

	expectTime := time.Now()
	expectedMovie := &domain.MovieWithoudID{
		Title:       "Test Movie",
		Description: "Test Description",
		ReleaseDate: expectTime,
		Rating:      5,
	}

	// OK. Search by title
	mock.ExpectQuery("SELECT movie_title, description, release_date, rating FROM Movies WHERE LOWER(movie_title) LIKE '%' || \\$1 || '%' LIMIT 1").
		WithArgs(strings.ToLower(dto1.Title)).
		WillReturnRows(sqlmock.NewRows([]string{"movie_title", "description", "release_date", "rating"}).
			AddRow("Test Movie", "Test Description", expectTime, 5))

	movie, err := storage.Get(dto1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if movie == nil {
		t.Error("expected movie, got nil")
	}

	if !reflect.DeepEqual(expectedMovie, movie) {
		t.Errorf("expected: %v, got: %v", expectedMovie, movie)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// OK. Search by actor
	dto2 := &domain.GetMovie{
		ActorName:  "Test Actor",
		SearchType: "actor",
	}
	mock.ExpectQuery("SELECT m.movie_title, m.description, m.release_date, m.rating FROM Movies m JOIN MovieActors ma ON m.movie_id = ma.movie_id JOIN Actors a ON ma.actor_id = a.actor_id WHERE LOWER(a.actor_name) LIKE '%' || \\$1 || '%' LIMIT 1").
		WithArgs(strings.ToLower(dto2.ActorName)).
		WillReturnRows(sqlmock.NewRows([]string{"movie_title", "description", "release_date", "rating"}).
			AddRow("Test Movie", "Test Description", expectTime, 5))

	movie, err = storage.Get(dto2)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if movie == nil {
		t.Error("expected movie, got nil")
	}

	if !reflect.DeepEqual(expectedMovie, movie) {
		t.Errorf("expected: %v, got: %v", expectedMovie, movie)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Invalid search type
	dto3 := &domain.GetMovie{
		SearchType: "invalid",
	}

	movie, err = storage.Get(dto3)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err.Error() != "invalid search type" {
		t.Errorf("expected error 'invalid search type', got '%s'", err.Error())
	}

	if movie != nil {
		t.Errorf("expected nil, got: %v", movie)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres returned error
	mock.ExpectQuery("SELECT m.movie_title, m.description, m.release_date, m.rating FROM Movies m JOIN MovieActors ma ON m.movie_id = ma.movie_id JOIN Actors a ON ma.actor_id = a.actor_id WHERE LOWER(a.actor_name) LIKE '%' || \\$1 || '%' LIMIT 1").
		WithArgs(strings.ToLower(dto2.ActorName)).
		WillReturnError(domain.ErrTest)

	movie, err = storage.Get(dto2)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if movie != nil {
		t.Errorf("expected nil, got: %v", movie)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMovieGetOrderedList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	storage := NewMovieStorage(db)

	dto := &domain.GetOrderedMovie{
		Order: "rating",
	}

	expectTime := time.Now()
	expectedMovies := []domain.MovieWithoudID{
		{Title: "Movie 1", Description: "Description 1", ReleaseDate: expectTime, Rating: 5},
		{Title: "Movie 2", Description: "Description 2", ReleaseDate: expectTime, Rating: 4},
	}

	// OK
	mock.ExpectQuery("SELECT movie_title, description, release_date, rating FROM Movies ORDER BY rating DESC").
		WillReturnRows(sqlmock.NewRows([]string{"movie_title", "description", "release_date", "rating"}).
			AddRow("Movie 1", "Description 1", expectTime, 5).
			AddRow("Movie 2", "Description 2", expectTime, 4))

	movies, err := storage.GetOrderedList(dto)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if !reflect.DeepEqual(expectedMovies, movies) {
		t.Errorf("expected: %v, got: %v", expectedMovies, movies)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Postgres returned error
	mock.ExpectQuery("SELECT movie_title, description, release_date, rating FROM Movies ORDER BY rating DESC").
		WillReturnError(domain.ErrTest)

	movies, err = storage.GetOrderedList(dto)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if movies != nil {
		t.Errorf("expected nil, got: %v", movies)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Rows scan error
	mock.ExpectQuery("SELECT movie_title, description, release_date, rating FROM Movies ORDER BY rating DESC").
		WillReturnRows(sqlmock.NewRows([]string{"movie_title", "description"}).
			AddRow("Movie 1", "Description 1"))

	movies, err = storage.GetOrderedList(dto)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if movies != nil {
		t.Errorf("expected nil, got: %v", movies)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
