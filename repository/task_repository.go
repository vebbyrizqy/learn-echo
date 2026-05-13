package repository

import (
	"learn-echo/config"
	"learn-echo/model"
)

func GetTasks() ([]model.Task, error) {

	var tasks []model.Task

	err := config.DB.
		Preload("User").
		Find(&tasks).Error

	return tasks, err
}

func GetTaskByID(id string) (model.Task, error) {

	var task model.Task

	err := config.DB.
		Preload("User").
		First(&task, id).Error

	return task, err
}

func CreateTask(task *model.Task) error {

	if err := config.DB.Create(task).Error; err != nil {
		return err
	}

	return config.DB.
		Preload("User").
		First(task, task.ID).Error
}

func UpdateTask(task *model.Task) error {

	if err := config.DB.Save(task).Error; err != nil {
		return err
	}

	return config.DB.
		Preload("User").
		First(task, task.ID).Error
}

func DeleteTask(task *model.Task) error {

	return config.DB.Delete(task).Error
}