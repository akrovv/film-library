package postgresqldb

import (
	"database/sql"

	"github.com/akrovv/filmlibrary/internal/domain"
)

type actorStorage struct {
	db *sql.DB
}

func NewActorStorage(db *sql.DB) *actorStorage {
	return &actorStorage{
		db: db,
	}
}

func (s *actorStorage) Create(dto *domain.CreateActor) error {
	_, err := s.db.Exec("INSERT INTO Actors (actor_name, gender, date_of_birth) VALUES ($1, $2, $3)", dto.Name,
		dto.Gender, dto.DateBirth)

	if err != nil {
		return err
	}

	return nil
}

func (s *actorStorage) Update(dto *domain.Actor) error {
	_, err := s.db.Exec("UPDATE Actors SET actor_name=$1, gender=$2, date_of_birth=$3 WHERE actor_id=$4", dto.Name, dto.Gender, dto.DateBirth, dto.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *actorStorage) Delete(dto *domain.DeleteActor) error {
	_, err := s.db.Exec("DELETE FROM Actors WHERE actor_id=$1", dto.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *actorStorage) GetList() ([]domain.ActorWithMovie, error) {
	actors := make([]domain.ActorWithMovie, 0, 4)

	rows, err := s.db.Query("SELECT actor_name, gender, date_of_birth, movie_title FROM Actors JOIN MovieActors USING(actor_id) JOIN Movies USING(movie_id)")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		actor := domain.ActorWithMovie{}

		err = rows.Scan(&actor.Name, &actor.Gender, &actor.DateBirth, &actor.Title)
		if err != nil {
			return nil, err
		}

		actors = append(actors, actor)
	}

	return actors, nil
}
