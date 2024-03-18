package domain

type User struct {
	Username string
	IsAdmin  bool
}

type CRUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserContext string
