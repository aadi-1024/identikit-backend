package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aadi-1024/identikit-backend/internal/database"
	"github.com/aadi-1024/identikit-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllSnippets(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := models.JsonResponse{}
		var data []models.Snippet
		var err error

		language := c.QueryParam("language")
		if language != "" {
			data, err = d.GetAllSnippets(c.Request().Context(), "language = ?", language)
		} else {
			data, err = d.GetAllSnippets(c.Request().Context())
		}

		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		res.Data = data
		res.Message = "successful"
		return c.JSON(http.StatusOK, res)
	}
}

func CreateSnippet(d *database.Database, v *validator.Validate) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := models.JsonResponse{}
		snippet := models.Snippet{}

		if err := json.NewDecoder(c.Request().Body).Decode(&snippet); err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		if err := v.Struct(snippet); err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		snippet.Id = uuid.New().String()

		if err := d.CreateSnippet(c.Request().Context(), snippet); err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, res)
		}

		res.Data = snippet
		res.Message = "successful"

		return c.JSON(http.StatusCreated, res)
	}
}
