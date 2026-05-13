package handler

import (
	"net/http"

	"learn-echo/config"
	"learn-echo/model"

	"github.com/labstack/echo/v5"
)

func GetTasks(c *echo.Context) error {

	var tasks []model.Task

	if err := config.DB.Find(&tasks).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *echo.Context) error {

	id := c.Param("id")

	var task model.Task

	if err := config.DB.First(&task, id).Error; err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "task not found",
		})
	}

	return c.JSON(http.StatusOK, task)
}

func CreateTask(c *echo.Context) error {

	t := new(model.Task)

	if err := c.Bind(t); err != nil {
		return err
	}

	if err := config.DB.Create(&t).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, t)
}

func UpdateTask(c *echo.Context) error {

	id := c.Param("id")

	var task model.Task

	if err := config.DB.First(&task, id).Error; err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "task not found",
		})
	}

	updatedTask := new(model.Task)

	if err := c.Bind(updatedTask); err != nil {
		return err
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}

	if updatedTask.UserID != 0 {
		task.UserID = updatedTask.UserID
	}

	task.Completed = updatedTask.Completed

	if err := config.DB.Save(&task).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func CompleteTask(c *echo.Context) error {

	id := c.Param("id")

	var task model.Task

	if err := config.DB.First(&task, id).Error; err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "task not found",
		})
	}

	task.Completed = true

	if err := config.DB.Save(&task).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c *echo.Context) error {

	id := c.Param("id")

	var task model.Task

	if err := config.DB.First(&task, id).Error; err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "task not found",
		})
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "task deleted",
	})
}