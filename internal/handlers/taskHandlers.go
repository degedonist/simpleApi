package handlers

import (
	"firstCoursePractice/internal/taskService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type taskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(service taskService.TaskService) *taskHandler {
	return &taskHandler{service: service}
}

func (h *taskHandler) GetTasks(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not get tasks")
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) AddTask(c echo.Context) error {
	var req taskService.RequestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	task, err := h.service.CreateTask(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not create task")
	}

	return c.JSON(http.StatusOK, task)
}

func (h *taskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")

	var req taskService.RequestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	updatedTask, err := h.service.UpdateTask(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not update task")
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *taskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	err := h.service.DeleteTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not delete task")
	}

	return c.NoContent(http.StatusNoContent)
}
