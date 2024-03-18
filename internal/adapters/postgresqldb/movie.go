package postgresqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/akrovv/filmlibrary/internal/domain"
)

type movieStorage struct {
	db *sql.DB
}

func NewMovieStorage(db *sql.DB) *movieStorage {
	return &movieStorage{
		db: db,
	}
}

func getSqlForMovieActors(actors []int64, lastID int64) (string, []interface{}, error) {
	var queryBuilder strings.Builder

	params := make([]interface{}, 0)

	_, err := queryBuilder.WriteString("INSERT INTO MovieActors (movie_id, actor_id) VALUES ")
	if err != nil {
		return "", nil, err
	}

	for index, actorID := range actors {
		params = append(params, lastID, actorID)
		if index > 0 {
			_, err = queryBuilder.WriteString(", ")
			if err != nil {
				return "", nil, err
			}
		}
		_, err = queryBuilder.WriteString(fmt.Sprintf("($%d, $%d)", 2*index+1, 2*index+2))
		if err != nil {
			return "", nil, err
		}
	}

	return queryBuilder.String(), params, nil
}

func (s *movieStorage) Create(dto *domain.CreateMovie) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var lastID int64
	err = tx.QueryRow("INSERT INTO Movies (movie_title, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING movie_id", dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating).Scan(&lastID)
	if err != nil {
		return err
	}

	cmd, params, err := getSqlForMovieActors(dto.Actors, lastID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(cmd, params...)
	if err != nil {
		return err
	}

	return nil
}

func (s *movieStorage) Update(dto *domain.Movie) error {
	tx, err := s.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}

		err = tx.Commit()
	}()

	_, err = tx.Exec("UPDATE Movies SET movie_title = $1, description = $2, release_date = $3, rating = $4 WHERE movie_id = $5",
		dto.Title, dto.Description, dto.ReleaseDate.Format("2006-01-02"), dto.Rating, dto.ID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM MovieActors WHERE movie_id = $1", dto.ID)
	if err != nil {
		return err
	}

	cmd, params, err := getSqlForMovieActors(dto.Actors, dto.ID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(cmd, params...)
	if err != nil {
		return err
	}

	return nil
}

func (s *movieStorage) Delete(dto *domain.DeleteMovie) error {
	_, err := s.db.Exec("DELETE FROM Movies WHERE movie_id = $1", dto.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *movieStorage) Get(dto *domain.GetMovie) (*domain.MovieWithoudID, error) {
	var query string
	var searchTerm string

	switch dto.SearchType {
	case "title":
		query = "SELECT movie_title, description, release_date, rating FROM Movies WHERE LOWER(movie_title) LIKE '%' || $1 || '%' LIMIT 1;"
		searchTerm = strings.ToLower(dto.Title)
	case "actor":
		query = "SELECT m.movie_title, m.description, m.release_date, m.rating FROM Movies m JOIN MovieActors ma ON m.movie_id = ma.movie_id JOIN Actors a ON ma.actor_id = a.actor_id WHERE LOWER(a.actor_name) LIKE '%' || $1 || '%' LIMIT 1;"
		searchTerm = strings.ToLower(dto.ActorName)
	default:
		return nil, errors.New("invalid search type")
	}

	movie := domain.MovieWithoudID{}
	if err := s.db.QueryRow(query, searchTerm).Scan(&movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating); err != nil {
		return nil, err
	}

	return &movie, nil
}

func (s *movieStorage) GetOrderedList(dto *domain.GetOrderedMovie) ([]domain.MovieWithoudID, error) {
	movies := make([]domain.MovieWithoudID, 0, 4)

	query := fmt.Sprintf("SELECT movie_title, description, release_date, rating FROM Movies ORDER BY %s DESC", dto.Order)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movie domain.MovieWithoudID
	for rows.Next() {
		err = rows.Scan(&movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
