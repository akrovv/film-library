package domain

import "time"

type Movie struct {
	ID          int64     `json:"movie_id"`
	Title       string    `json:"movie_title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      uint8     `json:"rating"`
	Actors      []int64   `json:"actors"`
}

type MovieWithoudID struct {
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      uint8
}

type CreateMovie struct {
	Title       string    `json:"movie_title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      uint8     `json:"rating"`
	Actors      []int64   `json:"actors"`
}

type DeleteMovie struct {
	ID int64 `json:"movie_id"`
}

type GetMovie struct {
	Title      string
	ActorName  string
	SearchType string
}

type GetOrderedMovie struct {
	Order string
}
