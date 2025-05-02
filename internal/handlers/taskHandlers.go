package handlers

import (
	"context"
	"firstCoursePractice/internal/models"
	"firstCoursePractice/internal/taskService"
	"firstCoursePractice/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"net/http"
)

type taskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(service taskService.TaskService) *taskHandler {
	return &taskHandler{service: service}
}

func (h *taskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, receivedTask := range allTasks {
		task := tasks.Task{
			Id:     &receivedTask.ID,
			Task:   &receivedTask.Task,
			IsDone: &receivedTask.IsDone,
		}

		response = append(response, task)
	}

	return response, nil
}

func (h *taskHandler) AddTask(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	req := request.Body

	taskToCreate := models.RequestBody{
		Task:   *req.Task,
		IsDone: *req.IsDone,
	}

	task, err := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}

	return response, nil
}

func (h *taskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")

	var req models.RequestBody
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
