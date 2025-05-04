package main

import (
	"firstCoursePractice/internal/db"
	"firstCoursePractice/internal/handlers"
	"firstCoursePractice/internal/repository"
	"firstCoursePractice/internal/taskService"
	"firstCoursePractice/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	taskRepo := repository.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	startErr := e.Start("localhost:8080")
	if startErr != nil {
		log.Fatalf("Could not start server: %v", startErr)
	}
}
