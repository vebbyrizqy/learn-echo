package service

import (
	"learn-echo/model"
	"learn-echo/repository"
)

func GetTasks() ([]model.Task, error) {

	return repository.GetTasks()
}

func GetTaskByID(id string) (model.Task, error) {

	return repository.GetTaskByID(id)
}

func CreateTask(task *model.Task) error {

	return repository.CreateTask(task)
}

func UpdateTask(id string, updatedTask *model.Task) (model.Task, error) {

	task, err := repository.GetTaskByID(id)
	if err != nil {
		return task, err
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}

	if updatedTask.UserID != 0 {
		task.UserID = updatedTask.UserID
	}

	task.Completed = updatedTask.Completed

	err = repository.UpdateTask(&task)

	return task, err
}

func CompleteTask(id string) (model.Task, error) {

	task, err := repository.GetTaskByID(id)
	if err != nil {
		return task, err
	}

	task.Completed = true

	err = repository.UpdateTask(&task)

	return task, err
}

func DeleteTask(id string) error {

	task, err := repository.GetTaskByID(id)
	if err != nil {
		return err
	}

	return repository.DeleteTask(&task)
}