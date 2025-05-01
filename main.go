package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var tasks []requestBody

type requestBody struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

func postHandler(c echo.Context) error {
	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task")
	}
	req.ID = uuid.New().String()
	tasks = append(tasks, req)
	return c.JSON(http.StatusCreated, "Task added")
}

func getHandler(c echo.Context) error {
	if len(tasks) != 0 {
		return c.JSON(http.StatusOK, tasks)
	}
	return c.JSON(http.StatusNotFound, "Task not found")
}

func patchHandler(c echo.Context) error {
	id := c.Param("id")
	var req requestBody

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Task = req.Task
			return c.JSON(http.StatusCreated, "Task updated")
		}
	}

	return c.JSON(http.StatusNotFound, "Task does not exist")
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, "Task does not exist")
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getHandler)
	e.POST("/tasks", postHandler)
	e.PATCH("/tasks/:id", patchHandler)
	e.DELETE("/tasks/:id", deleteHandler)

	e.Start("localhost:8080")
}
