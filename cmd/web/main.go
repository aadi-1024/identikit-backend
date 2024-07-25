package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	SetupRoutes(e)

	if err := e.Start("localhost:8080"); err != nil {
		log.Println(err.Error())
	}
}

