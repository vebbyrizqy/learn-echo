package helper

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func SuccessResponse(
	c *echo.Context,
	status int,
	message string,
	data interface{},
) error {

	return c.JSON(status, map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(
	c *echo.Context,
	status int,
	message string,
) error {

	return c.JSON(status, map[string]interface{}{
		"success": false,
		"message": message,
	})
}

func InternalServerError(c *echo.Context) error {

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"success": false,
		"message": "internal server error",
	})
}