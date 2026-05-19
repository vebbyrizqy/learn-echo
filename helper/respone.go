package helper

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func SuccessResponse(
	c *echo.Context,
	status int,
	message string,
	data interface{},
) error {

	response := APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	return c.JSON(status, response)
}

func ErrorResponse(
	c *echo.Context,
	status int,
	message string,
) error {

	response := APIResponse{
		Success: false,
		Message: message,
	}

	return c.JSON(status, response)
}

func InternalServerError(c *echo.Context) error {

	response := APIResponse{
		Success: false,
		Message: "internal server error",
	}

	return c.JSON(
		http.StatusInternalServerError,
		response,
	)
}