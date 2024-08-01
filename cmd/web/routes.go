package main

import (
	"net/http"

	"github.com/aadi-1024/identikit-backend/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	//user registration
	e.POST("/register", handlers.RegisterUserHandler(app.Db, app.Validate))
	e.POST("/login/dashboard", handlers.Login(app.Db, app.Validate, app.JwtSecret, "dashboard"))
	e.POST("/login/editor", handlers.Login(app.Db, app.Validate, app.JwtSecret, "editor"))

	snippets := e.Group("/snippets")
	snippets.GET("", handlers.GetAllSnippets(app.Db))
	snippets.POST("", handlers.CreateSnippet(app.Db, app.Validate))
	snippets.POST("/docs/:id", handlers.GenerateDocumentation(app.Db, app.GenaiClient))
	snippets.POST("/security/:id", handlers.GenerateSecurityAnalysis(app.Db, app.GenaiClient))
}
