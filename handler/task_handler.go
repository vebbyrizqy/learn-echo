package handler

import (
	"net/http"

	"learn-echo/helper"
	"learn-echo/model"
	"learn-echo/service"

	"github.com/labstack/echo/v5"
)

// GetTasks godoc
//
//	@Summary		Get all tasks
//	@Description	get all tasks
//	@Tags			tasks
//	@Produce		json
//	@Success		200	{object}	helper.APIResponse
//	@Router			/tasks [get]
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

// GetTaskByID godoc
//
//	@Summary		Get task by ID
//	@Description	get task by id
//	@Tags			tasks
//	@Produce		json
//	@Param			id	path		int	true	"Task ID"
//	@Success		200	{object}	helper.APIResponse
//	@Failure		404	{object}	helper.APIResponse
//	@Router			/tasks/{id} [get]
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

// CreateTask godoc
//
//	@Summary		Create task
//	@Description	create new task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.Task	true	"Task Request"
//	@Success		201		{object}	helper.APIResponse
//	@Failure		400		{object}	helper.APIResponse
//	@Router			/tasks [post]
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

// UpdateTask godoc
//
//	@Summary		Update task
//	@Description	update existing task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"Task ID"
//	@Param			request	body		model.Task	true	"Task Request"
//	@Success		200		{object}	helper.APIResponse
//	@Failure		404		{object}	helper.APIResponse
//	@Router			/tasks/{id} [patch]
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

// CompleteTask godoc
//
//	@Summary		Complete task
//	@Description	mark task as completed
//	@Tags			tasks
//	@Produce		json
//	@Param			id	path		int	true	"Task ID"
//	@Success		200	{object}	helper.APIResponse
//	@Failure		404	{object}	helper.APIResponse
//	@Router			/tasks/{id}/complete [patch]
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

// DeleteTask godoc
//
//	@Summary		Delete task
//	@Description	delete task by id
//	@Tags			tasks
//	@Produce		json
//	@Param			id	path		int	true	"Task ID"
//	@Success		200	{object}	helper.APIResponse
//	@Failure		404	{object}	helper.APIResponse
//	@Router			/tasks/{id} [delete]
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