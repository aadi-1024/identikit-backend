package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aadi-1024/identikit-backend/internal/database"
	"github.com/aadi-1024/identikit-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RegisterUserHandler(d *database.Database, v *validator.Validate) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}
		res := models.JsonResponse{}

		err := json.NewDecoder(c.Request().Body).Decode(&user)
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		err = v.Struct(user)
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		_, err = d.RegisterUser(c.Request().Context(), user)
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		res.Message = "successful"
		return c.JSON(http.StatusCreated, res)
	}
}

func createJwt(id int, role, scope string, jwtSecret []byte) (string, error) {
	clms := models.Claims{}

	clms.Role = role
	clms.Id = id
	clms.Scope = scope
	clms.ExpiresAt = jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clms)
	return token.SignedString(jwtSecret)
}

func Login(d *database.Database, v *validator.Validate, jwtSecret []byte, scope string) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}
		res := models.JsonResponse{}

		err := json.NewDecoder(c.Request().Body).Decode(&user)
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		err = v.StructExcept(user, "Role")
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		ret, err := d.LoginUser(c.Request().Context(), user.Email, user.Password)
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusUnauthorized, res)
		}

		token, err := createJwt(ret.Id, ret.Role, scope, jwtSecret)
		if err != nil {
			res.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, res)
		}

		res.Message = "successful"
		res.Data = token

		return c.JSON(http.StatusOK, res)
	}
}
