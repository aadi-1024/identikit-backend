package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aadi-1024/identikit-backend/internal/database"
	"github.com/aadi-1024/identikit-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/generative-ai-go/genai"
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

func GenerateDocumentation(d *database.Database, ai *genai.Client) echo.HandlerFunc {
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
		code := ""

		snippet, err := d.GetSnippetById(c.Request().Context(), id)
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		for _, v := range snippet.Body {
			str, ok := v.(string)
			if !ok {
				continue
			}
			code += str + "\n"
		}

		if forceGeneration || snippet.Documentation == "" {
			model := ai.GenerativeModel("gemini-1.5-flash")
			resp, err := model.GenerateContent(c.Request().Context(), genai.Text("Generate documentation for the following code snippet: "), genai.Text(code))
			if err != nil {
				res.Message = err.Error()
				return c.JSON(http.StatusInternalServerError, res)
			}

			for _, cand := range resp.Candidates {
				if cand.Content != nil {
					for _, part := range cand.Content.Parts {
						documentation += fmt.Sprint(part)
					}
				}
			}

			d.UpdateSnippet(c.Request().Context(), models.Snippet{Id: snippet.Id, Documentation: documentation})
		} else {
			documentation = snippet.Documentation
		}

		res.Message = "successful"
		res.Data = documentation
		return c.JSON(http.StatusOK, res)
	}
}

func GenerateSecurityAnalysis(d *database.Database, ai *genai.Client) echo.HandlerFunc {
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
		code := ""

		snippet, err := d.GetSnippetById(c.Request().Context(), id)
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		for _, v := range snippet.Body {
			str, ok := v.(string)
			if !ok {
				continue
			}
			code += str + "\n"
		}

		if forceGeneration || snippet.Security == "" {
			model := ai.GenerativeModel("gemini-1.5-flash")
			resp, err := model.GenerateContent(c.Request().Context(), genai.Text("Generate security report for the following code snippet: "), genai.Text(code))
			if err != nil {
				res.Message = err.Error()
				return c.JSON(http.StatusInternalServerError, res)
			}

			for _, cand := range resp.Candidates {
				if cand.Content != nil {
					for _, part := range cand.Content.Parts {
						documentation += fmt.Sprint(part)
					}
				}
			}

			d.UpdateSnippet(c.Request().Context(), models.Snippet{Id: snippet.Id, Security: documentation})
		} else {
			documentation = snippet.Security
		}

		res.Message = "successful"
		res.Data = documentation
		return c.JSON(http.StatusOK, res)
	}
}
