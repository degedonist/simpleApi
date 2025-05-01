package main

import (
	"firstCoursePractice/internal/db"
	"firstCoursePractice/internal/handlers"
	"firstCoursePractice/internal/taskService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandlers.GetTasks)
	e.POST("/tasks", taskHandlers.AddTask)
	e.PATCH("/tasks/:id", taskHandlers.UpdateTask)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTask)

	e.Start("localhost:8080")
}
