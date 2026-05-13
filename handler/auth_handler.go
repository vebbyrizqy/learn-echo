package handler

import (
	"net/http"

	"learn-echo/helper"
	"learn-echo/model"
	"learn-echo/service"

	"github.com/labstack/echo/v5"
)

func Login(c *echo.Context) error {

	req := new(model.LoginRequest)

	if err := c.Bind(req); err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
		)
	}

	token, err := service.Login(req)
	if err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusUnauthorized,
			"invalid credentials",
		)
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"login successful",
		map[string]string{
			"token": token,
		},
	)
}

func Profile(c *echo.Context) error {

	userID := c.Get("user_id")
	name := c.Get("name")
	email := c.Get("email")

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"profile fetched successfully",
		map[string]interface{}{
			"user_id": userID,
			"name":    name,
			"email":   email,
		},
	)
}