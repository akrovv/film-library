package domain

type CreateSession struct {
	Username string
	IsAdmin  bool
}

type GetSession struct {
	Username string
}
