package taskService

import (
	"errors"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(req TaskRequest) (Task, error)
	GetTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, req TaskRequest) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(req TaskRequest) (Task, error) {
	var task Task

	if req.Task == "" {
		return Task{}, errors.New("task can't be empty")
	}

	task.ID = uuid.New().String()
	task.Task = req.Task
	task.IsDone = req.IsDone
	task.UserID = req.UserID

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) GetTasks() ([]Task, error) {
	return s.repo.GetTasks()
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id string, req TaskRequest) (Task, error) {
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
