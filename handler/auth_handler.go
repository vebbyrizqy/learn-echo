package handler

import (
	"net/http"

	"learn-echo/data"
	"learn-echo/model"
	"learn-echo/utils"

	"github.com/labstack/echo/v5"
)

func Login(c *echo.Context) error {

	req := new(model.LoginRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	for _, user := range data.Users {

		if user.Email == req.Email && utils.CheckPasswordHash(req.Password, user.Password) {

			token, err := utils.GenerateJWT(user.ID, user.Email)
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, map[string]string{
				"token": token,
			})
		}
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "invalid credentials",
	})
}

func Profile(c *echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"message": "welcome to profile",
	})
}