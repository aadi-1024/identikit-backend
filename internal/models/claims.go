package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	Role  string `json:"role"`
	Id    int    `json:"id"`
	Scope string `json:"scope"`
}
