package taskService

import (
	"errors"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(req RequestBody) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, req RequestBody) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func (s *taskService) CreateTask(req RequestBody) (Task, error) {
	var task Task

	if req.Task == "" {
		return Task{}, errors.New("task can't be empty")
	}

	task.ID = uuid.New().String()
	task.Task = req.Task
	task.IsDone = req.IsDone

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id string, req RequestBody) (Task, error) {
	if req.Task == "" {
		return Task{}, errors.New("task can't be empty")
	}

	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	task.Task = req.Task
	task.IsDone = req.IsDone

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}
