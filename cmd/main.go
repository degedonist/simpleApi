package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func initDatabase() {
	dsn := "host=localhost user=postgres password=secretpass dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

type Task struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func postHandler(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task")
	}
	req.ID = uuid.New().String()

	if err := db.Create(&req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to add task")
	}

	return c.JSON(http.StatusCreated, "Task added")
}

func getHandler(c echo.Context) error {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get tasks")
	}

	if len(tasks) == 0 {
		return c.JSON(http.StatusNotFound, "No tasks found")
	}

	return c.JSON(http.StatusOK, tasks)
}

func patchHandler(c echo.Context) error {
	id := c.Param("id")

	var req RequestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task")
	}

	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Task not found")
	}

	task.Task = req.Task
	task.IsDone = req.IsDone

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update task")
	}

	return c.JSON(http.StatusOK, task)
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete task")
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDatabase()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getHandler)
	e.POST("/tasks", postHandler)
	e.PATCH("/tasks/:id", patchHandler)
	e.DELETE("/tasks/:id", deleteHandler)

	e.Start("localhost:8080")
}
