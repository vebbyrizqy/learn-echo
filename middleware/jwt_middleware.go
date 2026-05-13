package middleware

import (
	"net/http"
	"strings"

	"learn-echo/helper"
	"learn-echo/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c *echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {

			return helper.ErrorResponse(
				c,
				http.StatusUnauthorized,
				"missing token",
			)
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utils.SECRET_KEY, nil
		})

		if err != nil || !token.Valid {

			return helper.ErrorResponse(
				c,
				http.StatusUnauthorized,
				"invalid token",
			)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {

			return helper.ErrorResponse(
				c,
				http.StatusUnauthorized,
				"invalid token claims",
			)
		}

		c.Set("user_id", claims["user_id"])
		c.Set("name", claims["name"])
		c.Set("email", claims["email"])

		return next(c)
	}
}