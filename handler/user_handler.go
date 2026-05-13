package handler

import (
	"net/http"
	"strconv"

	"learn-echo/data"
	"learn-echo/model"
	"learn-echo/utils"

	"github.com/labstack/echo/v5"
)

func GetUsers(c *echo.Context) error {
	return c.JSON(http.StatusOK, data.Users)
}

func GetUserByID(c *echo.Context) error {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
	}

	for _, user := range data.Users {

		if user.ID == userID {
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "user not found",
	})
}

func CreateUser(c *echo.Context) error {

	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	
	u.Password = hashedPassword
	
	u.ID = len(data.Users) + 1
	
	data.Users = append(data.Users, *u)

	return c.JSON(http.StatusCreated, u)
}

func UpdateUser(c *echo.Context) error {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
	}

	updatedUser := new(model.User)

	if err := c.Bind(updatedUser); err != nil {
		return err
	}

	for i, user := range data.Users {

		if user.ID == userID {

			if updatedUser.Name != "" {
				data.Users[i].Name = updatedUser.Name
			}

			if updatedUser.Email != "" {
				data.Users[i].Email = updatedUser.Email
			}

			if updatedUser.Password != "" {

				hashedPassword, err := utils.HashPassword(updatedUser.Password)
				if err != nil {
					return err
				}

				data.Users[i].Password = hashedPassword
			}

			return c.JSON(http.StatusOK, data.Users[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "user not found",
	})
}

func DeleteUser(c *echo.Context) error {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
	}

	for i, user := range data.Users {

		if user.ID == userID {

			data.Users = append(data.Users[:i], data.Users[i+1:]...)

			return c.JSON(http.StatusOK, map[string]string{
				"message": "user deleted",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "user not found",
	})
}