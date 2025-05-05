package handlers

import (
	"context"
	"firstCoursePractice/internal/taskService"
	"firstCoursePractice/internal/web/tasks"
)

type taskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(service taskService.TaskService) *taskHandler {
	return &taskHandler{service: service}
}

// Get all existing tasks
func (h *taskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, receivedTask := range allTasks {
		task := tasks.Task{
			Id:     &receivedTask.ID,
			Task:   &receivedTask.Task,
			IsDone: &receivedTask.IsDone,
			UserId: &receivedTask.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

// Create a new task
func (h *taskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	req := request.Body

	taskToCreate := taskService.TaskRequest{
		Task:   *req.Task,
		IsDone: *req.IsDone,
		UserID: *req.UserId,
	}

	task, err := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		IsDone: &task.IsDone,
		UserId: &task.UserID,
	}

	return response, nil
}

// Update task by ID
func (h *taskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	req := request.Body

	taskToUpdate := taskService.TaskRequest{
		Task:   *req.Task,
		IsDone: *req.IsDone,
	}

	updatedTask, err := h.service.UpdateTask(id, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

// Delete task by ID
func (h *taskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	err := h.service.DeleteTask(id)
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId204Response{}

	return response, nil
}
