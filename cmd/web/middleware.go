package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/aadi-1024/identikit-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(jwtSecret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := models.JsonResponse{}
			authHeader := c.Request().Header.Get("Authorization")
			spl := strings.Split(authHeader, " ")

			if !strings.HasPrefix(authHeader, "bearer") || len(spl) != 2 {
				res.Message = "malformed Authorization header"
				return c.JSON(http.StatusUnauthorized, res)
			}

			clms := models.Claims{}
			token := spl[1]
			tkn, err := jwt.ParseWithClaims(token, clms, func(t *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})

			if !tkn.Valid {
				res.Message = "invalid token"
				return c.JSON(http.StatusUnauthorized, res)
			}

			exp, err := tkn.Claims.GetExpirationTime()
			if err != nil {
				res.Message = err.Error()
				return c.JSON(http.StatusInternalServerError, res)
			}

			if exp.Before(time.Now()) {
				res.Message = "session expired"
				return c.JSON(http.StatusUnauthorized, res)
			}
			return next(c)
		}
	}
}
