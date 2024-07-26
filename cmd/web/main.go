package main

import (
	"log"
	"time"

	"github.com/aadi-1024/identikit-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var app App

func main() {
	app = App{}
	e := echo.New()

	db, err := database.InitDb("postgres://postgres:password@localhost:5432/identikit", 3*time.Second)
	if err != nil {
		log.Fatalln(err.Error())
	}
	app.Db = db
	app.JwtSecret = []byte("HUGE_SECRET")

	validate := validator.New(validator.WithRequiredStructEnabled())
	app.Validate = validate

	SetupRoutes(e)

	if err := e.Start("localhost:8080"); err != nil {
		log.Println(err.Error())
	}
}
