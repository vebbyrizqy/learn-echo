package repository

import (
	"learn-echo/config"
	"learn-echo/model"
)

func GetUsers() ([]model.User, error) {

	var users []model.User

	err := config.DB.Find(&users).Error

	return users, err
}

func GetUserByID(id string) (model.User, error) {

	var user model.User

	err := config.DB.First(&user, id).Error

	return user, err
}

func CreateUser(user *model.User) error {

	return config.DB.Create(user).Error
}

func UpdateUser(user *model.User) error {

	return config.DB.Save(user).Error
}

func DeleteUser(user *model.User) error {

	return config.DB.Delete(user).Error
}