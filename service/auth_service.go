package service

import (
	"errors"

	"learn-echo/model"
	"learn-echo/repository"
	"learn-echo/utils"
)

func Login(req *model.LoginRequest) (model.User, error) {

	user, err := repository.FindUserByEmail(req.Email)

	if err != nil {

		return model.User{}, errors.New("invalid credentials")
	}

	isValidPassword := utils.CheckPasswordHash(
		req.Password,
		user.Password,
	)

	if !isValidPassword {

		return model.User{}, errors.New("invalid credentials")
	}

	return user, nil
}