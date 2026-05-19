package handler

import (
	"net/http"
	"time"

	"learn-echo/helper"
	"learn-echo/model"
	"learn-echo/service"
	"learn-echo/utils"

	"github.com/labstack/echo/v5"
)

// Login godoc
//
//	@Summary		Login user
//	@Description	login using email and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.LoginRequest	true	"Login Request"
//	@Success		200		{object}	helper.APIResponse
//	@Failure		401		{object}	helper.APIResponse
//	@Router			/login [post]
func Login(c *echo.Context) error {

	req := new(model.LoginRequest)

	if err := c.Bind(req); err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
		)
	}

	user, err := service.Login(req)

	if err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusUnauthorized,
			"invalid credentials",
		)
	}

	sessionID := utils.GenerateSessionID()

	utils.Sessions[sessionID] = map[string]interface{}{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
	}

	cookie := new(http.Cookie)

	cookie.Name = "session_id"
	cookie.Value = sessionID

	cookie.Expires = time.Now().Add(24 * time.Hour)

	cookie.HttpOnly = true
	cookie.Path = "/"

	cookie.Secure = false

	cookie.SameSite = http.SameSiteLaxMode

	c.SetCookie(cookie)

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"login successful",
		nil,
	)
}

// Profile godoc
//
//	@Summary		Get user profile
//	@Description	get logged in user profile
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	helper.APIResponse
//	@Failure		401	{object}	helper.APIResponse
//	@Router			/profile [get]
func Profile(c *echo.Context) error {

	user := model.User{
		ID:    c.Get("user_id").(int),
		Name:  c.Get("name").(string),
		Email: c.Get("email").(string),
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"profile fetched successfully",
		user,
	)
}

// Logout godoc
//
//	@Summary		Logout user
//	@Description	logout current session
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	helper.APIResponse
//	@Router			/logout [post]
func Logout(c *echo.Context) error {

	cookie, err := c.Cookie("session_id")

	if err == nil {

		delete(utils.Sessions, cookie.Value)
	}

	expiredCookie := new(http.Cookie)

	expiredCookie.Name = "session_id"
	expiredCookie.Value = ""
	expiredCookie.Path = "/"
	expiredCookie.MaxAge = -1

	c.SetCookie(expiredCookie)

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"logout successful",
		nil,
	)
}