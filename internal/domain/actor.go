package domain

import "time"

type Actor struct {
	ID        int64     `json:"actor_id"`
	Name      string    `json:"actor_name"`
	Gender    string    `json:"gender"`
	DateBirth time.Time `json:"date_of_birth"`
}

type CreateActor struct {
	Name      string    `json:"actor_name"`
	Gender    string    `json:"gender"`
	DateBirth time.Time `json:"date_of_birth"`
}

type DeleteActor struct {
	ID int64 `json:"actor_id"`
}

type ActorWithMovie struct {
	CreateActor
	Title string
}
