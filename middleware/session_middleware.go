package middleware

import (
	"net/http"

	"learn-echo/helper"
	"learn-echo/utils"

	"github.com/labstack/echo/v5"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c *echo.Context) error {

		cookie, err := c.Cookie("session_id")

		if err != nil {

			return helper.ErrorResponse(
				c,
				http.StatusUnauthorized,
				"unauthorized",
			)
		}

		session, exists := utils.Sessions[cookie.Value]

		if !exists {

			return helper.ErrorResponse(
				c,
				http.StatusUnauthorized,
				"invalid session",
			)
		}

		c.Set("user_id", session["user_id"])
		c.Set("name", session["name"])
		c.Set("email", session["email"])

		return next(c)
	}
}