package taskService

import (
	"errors"
	"firstCoursePractice/internal/models"
	"firstCoursePractice/internal/repository"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(req models.RequestBody) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id string) (models.Task, error)
	UpdateTask(id string, req models.RequestBody) (models.Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo repository.TaskRepository
}

func (s *taskService) CreateTask(req models.RequestBody) (models.Task, error) {
	var task models.Task

	if req.Task == "" {
		return models.Task{}, errors.New("task can't be empty")
	}

	task.ID = uuid.New().String()
	task.Task = req.Task
	task.IsDone = req.IsDone

	if err := s.repo.CreateTask(task); err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (models.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id string, req models.RequestBody) (models.Task, error) {
	if req.Task == "" {
		return models.Task{}, errors.New("task can't be empty")
	}

	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return models.Task{}, err
	}

	task.Task = req.Task
	task.IsDone = req.IsDone

	if err := s.repo.UpdateTask(task); err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{repo: r}
}
