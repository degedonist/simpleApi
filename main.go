package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func postHandler(c echo.Context) error {
	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task")
	}

	task = req.Task
	return c.JSON(http.StatusCreated, "Task added")
}

func getHandler(c echo.Context) error {
	if task != "" {
		return c.JSON(http.StatusOK, task)
	}
	return c.JSON(http.StatusNotFound, "Task not found")
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", getHandler)
	e.POST("/", postHandler)

	e.Start("localhost:8080")
}
