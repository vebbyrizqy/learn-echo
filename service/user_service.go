package service

import (
	"learn-echo/model"
	"learn-echo/repository"
	"learn-echo/utils"
)

func GetUsers() ([]model.User, error) {

	return repository.GetUsers()
}

func GetUserByID(id string) (model.User, error) {

	return repository.GetUserByID(id)
}

func CreateUser(user *model.User) error {

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return repository.CreateUser(user)
}

func UpdateUser(id string, updatedUser *model.User) (model.User, error) {

	user, err := repository.GetUserByID(id)
	if err != nil {
		return user, err
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
			return user, err
		}

		user.Password = hashedPassword
	}

	err = repository.UpdateUser(&user)

	return user, err
}

func DeleteUser(id string) error {

	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}

	return repository.DeleteUser(&user)
}