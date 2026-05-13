package handler

import (
	"net/http"

	"learn-echo/helper"
	"learn-echo/model"
	"learn-echo/service"

	"github.com/labstack/echo/v5"
)

func GetTasks(c *echo.Context) error {

	tasks, err := service.GetTasks()
	if err != nil {
		return helper.InternalServerError(c)
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"tasks fetched successfully",
		tasks,
	)
}

func GetTaskByID(c *echo.Context) error {

	id := c.Param("id")

	task, err := service.GetTaskByID(id)
	if err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusNotFound,
			"task not found",
		)
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"task fetched successfully",
		task,
	)
}

func CreateTask(c *echo.Context) error {

	t := new(model.Task)

	if err := c.Bind(t); err != nil {
		return helper.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
		)
	}

	if err := service.CreateTask(t); err != nil {
		return helper.InternalServerError(c)
	}

	return helper.SuccessResponse(
		c,
		http.StatusCreated,
		"task created successfully",
		t,
	)
}

func UpdateTask(c *echo.Context) error {

	id := c.Param("id")

	updatedTask := new(model.Task)

	if err := c.Bind(updatedTask); err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
		)
	}

	task, err := service.UpdateTask(id, updatedTask)
	if err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusNotFound,
			"task not found",
		)
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"task updated successfully",
		task,
	)
}

func CompleteTask(c *echo.Context) error {

	id := c.Param("id")

	task, err := service.CompleteTask(id)
	if err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusNotFound,
			"task not found",
		)
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"task completed successfully",
		task,
	)
}

func DeleteTask(c *echo.Context) error {

	id := c.Param("id")

	if err := service.DeleteTask(id); err != nil {

		return helper.ErrorResponse(
			c,
			http.StatusNotFound,
			"task not found",
		)
	}

	return helper.SuccessResponse(
		c,
		http.StatusOK,
		"task deleted successfully",
		nil,
	)
}