package models

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,lt=72,gt=8"`
	Role     string `json:"role" validate:"required,oneof=user admin"`
}
