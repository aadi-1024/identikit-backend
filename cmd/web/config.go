package main

import (
	"github.com/aadi-1024/identikit-backend/internal/database"
	"github.com/go-playground/validator/v10"
)

type App struct {
	Db        *database.Database
	JwtSecret []byte
	Validate  *validator.Validate
}
