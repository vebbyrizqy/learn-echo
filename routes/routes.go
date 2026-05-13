package routes

import (
	"learn-echo/handler"

	"learn-echo/middleware"

	"github.com/labstack/echo/v5"
)


func InitRoutes(e *echo.Echo) {

	// user routes
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUserByID)

	e.POST("/users", handler.CreateUser)

	e.PATCH("/users/:id", handler.UpdateUser)

	e.DELETE("/users/:id", handler.DeleteUser)

	// task routes
	e.GET("/tasks", handler.GetTasks)
	e.GET("/tasks/:id", handler.GetTaskByID)

	e.POST("/tasks", handler.CreateTask)

	e.PUT("/tasks/:id", handler.UpdateTask)
	e.PUT("/tasks/:id/complete", handler.CompleteTask)

	e.DELETE("/tasks/:id", handler.DeleteTask)

	// auth routes
	e.POST("/login", handler.Login)
	e.GET("/profile", handler.Profile, middleware.JWTMiddleware)
}