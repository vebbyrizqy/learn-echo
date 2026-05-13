package handler

import (
	"net/http"
	"strconv"

	"learn-echo/data"
	"learn-echo/model"

	"github.com/labstack/echo/v5"
)

func GetTasks(c *echo.Context) error {
	return c.JSON(http.StatusOK, data.Tasks)
}

func GetTaskByID(c *echo.Context) error {

	id := c.Param("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
	}

	for _, task := range data.Tasks {

		if task.ID == taskID {
			return c.JSON(http.StatusOK, task)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "task not found",
	})
}

func CreateTask(c *echo.Context) error {

	t := new(model.Task)

	if err := c.Bind(t); err != nil {
		return err
	}

	t.ID = len(data.Tasks) + 1

	data.Tasks = append(data.Tasks, *t)

	return c.JSON(http.StatusCreated, t)
}

func UpdateTask(c *echo.Context) error {

	id := c.Param("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
	}

	updatedTask := new(model.Task)

	if err := c.Bind(updatedTask); err != nil {
		return err
	}

	for i, task := range data.Tasks {

		if task.ID == taskID {

			if updatedTask.Title != "" {
				data.Tasks[i].Title = updatedTask.Title
			}

			if updatedTask.UserID != 0 {
				data.Tasks[i].UserID = updatedTask.UserID
			}

			data.Tasks[i].Completed = updatedTask.Completed

			return c.JSON(http.StatusOK, data.Tasks[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "task not found",
	})
}
func CompleteTask(c *echo.Context) error {

	id := c.Param("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
	}

	for i, task := range data.Tasks {

		if task.ID == taskID {
			data.Tasks[i].Completed = true

			return c.JSON(http.StatusOK, data.Tasks[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "task not found",
	})
}

func DeleteTask(c *echo.Context) error {

	id := c.Param("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
	}

	for i, task := range data.Tasks {

		if task.ID == taskID {

			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)

			return c.JSON(http.StatusOK, map[string]string{
				"message": "task deleted",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "task not found",
	})
}