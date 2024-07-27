package handlers

import (
	"encoding/json"
	"net/http"
	"time"

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

func GenerateDocumentation(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		var forceGeneration bool
		res := models.JsonResponse{}
		if c.QueryParam("force") == "true" {
			forceGeneration = true
		} else {
			forceGeneration = false
		}

		id := c.Param("id")
		var documentation string

		if forceGeneration {
			documentation = "sample documentation at " + time.Now().Format(time.DateTime)
			d.UpdateSnippet(c.Request().Context(), models.Snippet{Id: id, Documentation: documentation})
		} else {
			snip, err := d.GetSnippetById(c.Request().Context(), id)
			if err != nil {
				res.Message = err.Error()
				return c.JSON(http.StatusBadRequest, res)
			}

			if snip.Documentation == "" {
				documentation = "sample documentation at " + time.Now().Format(time.DateTime)
				d.UpdateSnippet(c.Request().Context(), models.Snippet{Id: id, Documentation: documentation})
			} else {
				documentation = snip.Documentation
			}
		}
		res.Message = "successful"
		res.Data = documentation
		return c.JSON(http.StatusOK, res)
	}
}


func GenerateSecurityAnalysis(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		var forceGeneration bool
		res := models.JsonResponse{}
		if c.QueryParam("force") == "true" {
			forceGeneration = true
		} else {
			forceGeneration = false
		}

		id := c.Param("id")
		var profile string

		if forceGeneration {
			profile = "sample security profile at " + time.Now().Format(time.DateTime)
			d.UpdateSnippet(c.Request().Context(), models.Snippet{Id: id, Security: profile})
		} else {
			snip, err := d.GetSnippetById(c.Request().Context(), id)
			if err != nil {
				res.Message = err.Error()
				return c.JSON(http.StatusBadRequest, res)
			}

			if snip.Security == "" {
				profile = "sample security profile at " + time.Now().Format(time.DateTime)
				d.UpdateSnippet(c.Request().Context(), models.Snippet{Id: id, Security: profile})
			} else {
				profile = snip.Security
			}
		}
		res.Message = "successful"
		res.Data = profile
		return c.JSON(http.StatusOK, res)
	}
}
