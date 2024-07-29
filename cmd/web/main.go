package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aadi-1024/identikit-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

var app App

func main() {
	app = App{}
	e := echo.New()

	client, err := genai.NewClient(context.Background(), option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer client.Close()

	app.GenaiClient = client

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
