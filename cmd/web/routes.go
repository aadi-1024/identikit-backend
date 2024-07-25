package main

import (
	"net/http"

	"github.com/aadi-1024/identikit-backend/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	//user registration
	e.POST("/register", handlers.RegisterUserHandler(app.Db, app.Validate))
	e.POST("/login/dashboard", handlers.Login(app.Db, app.Validate, app.JwtSecret, "dashboard"))
	e.POST("/login/editor", handlers.Login(app.Db, app.Validate, app.JwtSecret, "editor"))
}
