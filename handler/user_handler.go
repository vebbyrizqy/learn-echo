package handler

import (
	"net/http"

	"learn-echo/config"
	"learn-echo/model"
	"learn-echo/utils"
	"learn-echo/helper"

	"github.com/labstack/echo/v5"
)

func GetUsers(c *echo.Context) error {

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return err
	}

	return helper.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", users)
}

func GetUserByID(c *echo.Context) error {

	id := c.Param("id")

	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {

		return helper.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	return helper.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
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

	return helper.SuccessResponse(c, http.StatusCreated, "User created successfully", u)
}

func UpdateUser(c *echo.Context) error {

	id := c.Param("id")

	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {

		return helper.ErrorResponse(c, http.StatusNotFound, "User not found")
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

	return helper.SuccessResponse(c, http.StatusOK, "User updated successfully", user)
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

	return helper.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}