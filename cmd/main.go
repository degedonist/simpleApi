package main

import (
	"firstCoursePractice/internal/db"
	"firstCoursePractice/internal/handlers"
	"firstCoursePractice/internal/taskService"
	"firstCoursePractice/internal/userService"
	"firstCoursePractice/internal/web/tasks"
	"firstCoursePractice/internal/web/users"
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
	taskHandler := handlers.NewTaskHandler(taskService)
	userRepo := userService.NewUserRepository(database)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)
	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	startErr := e.Start("localhost:8080")
	if startErr != nil {
		log.Fatalf("Could not start server: %v", startErr)
	}
}
