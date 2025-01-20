package model

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	AuthRequest
	Id   int
	Role string
}
