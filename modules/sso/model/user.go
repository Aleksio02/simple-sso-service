package model

type AuthRequest struct {
	Username string
	Password string
}

type User struct {
	AuthRequest
	Id   int
	Role string
}
