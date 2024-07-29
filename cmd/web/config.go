package main

import (
	"github.com/aadi-1024/identikit-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/generative-ai-go/genai"
)

type App struct {
	Db          *database.Database
	GenaiClient *genai.Client
	JwtSecret   []byte
	Validate    *validator.Validate
}
