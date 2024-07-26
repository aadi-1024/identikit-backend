package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required,lte=72,gte=8"`
	Role     string `json:"role" validate:"required,oneof=user admin" gorm:"default:user"`
}
