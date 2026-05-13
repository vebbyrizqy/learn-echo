package handler

import (
	"net/http"

	"learn-echo/config"
	"learn-echo/model"
	"learn-echo/utils"

	"github.com/labstack/echo/v5"
)

func GetUsers(c *echo.Context) error {

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserByID(c *echo.Context) error {

	id := c.Param("id")

	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "user not found",
		})
	}

	return c.JSON(http.StatusOK, user)
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
	
	if err := config.DB.Create(&u).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func UpdateUser(c *echo.Context) error {

	id := c.Param("id")

	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "user not found",
		})
	}

	updatedUser := new(model.User)

	if err := c.Bind(updatedUser); err != nil {
		return err
	}

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	if updatedUser.Password != "" {

		hashedPassword, err := utils.HashPassword(updatedUser.Password)
		if err != nil {
			return err
		}

		user.Password = hashedPassword
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c *echo.Context) error {

	id := c.Param("id")

	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "user not found",
		})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user deleted",
	})
}