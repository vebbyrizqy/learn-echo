package handler

import (
	"net/http"

	"learn-echo/config"
	"learn-echo/helper"
	"learn-echo/model"
	"learn-echo/utils"

	"github.com/labstack/echo/v5"
)

// GetUsers godoc
//
//	@Summary		Get all users
//	@Description	get all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	helper.APIResponse
//	@Router			/users [get]
func GetUsers(c *echo.Context) error {

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return err
	}

	return helper.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", users)
}

// GetUserByID godoc
//
//	@Summary		Get user by ID
//	@Description	get user by id
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	helper.APIResponse
//	@Failure		404	{object}	helper.APIResponse
//	@Router			/users/{id} [get]
func GetUserByID(c *echo.Context) error {

	id := c.Param("id")

	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {

		return helper.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	return helper.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

// CreateUser godoc
//
//	@Summary		Create user
//	@Description	create new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.User	true	"User Request"
//	@Success		201		{object}	helper.APIResponse
//	@Failure		400		{object}	helper.APIResponse
//	@Router			/users [post]
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

// UpdateUser godoc
//
//	@Summary		Update user
//	@Description	update existing user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"User ID"
//	@Param			request	body		model.User	true	"User Request"
//	@Success		200		{object}	helper.APIResponse
//	@Failure		404		{object}	helper.APIResponse
//	@Router			/users/{id} [patch]
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

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	delete user by id
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	helper.APIResponse
//	@Failure		404	{object}	helper.APIResponse
//	@Router			/users/{id} [delete]
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