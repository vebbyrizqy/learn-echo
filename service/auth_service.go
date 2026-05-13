package service

import (
	"errors"

	"learn-echo/model"
	"learn-echo/repository"
	"learn-echo/utils"
)

func Login(req *model.LoginRequest) (string, error) {

	user, err := repository.FindUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	isValidPassword := utils.CheckPasswordHash(
		req.Password,
		user.Password,
	)

	if !isValidPassword {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID,user.Name, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}