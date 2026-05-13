package repository

import (
	"learn-echo/config"
	"learn-echo/model"
)

func FindUserByEmail(email string) (model.User, error) {

	var user model.User

	err := config.DB.
		Where("email = ?", email).
		First(&user).Error

	return user, err
}