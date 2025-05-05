package userService

import (
	"errors"
	"firstCoursePractice/internal/taskService"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]User, error)
	GetUserById(id string) (User, error)
	GetTasksForUser(userID string) ([]taskService.Task, error)
	CreateUser(req UserRequest) (User, error)
	UpdateUser(id string, req UserRequest) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) GetAllUsers() ([]User, error) {
	return u.repo.GetAllUsers()
}

func (u *userService) GetUserById(id string) (User, error) {
	return u.repo.GetUserById(id)
}

func (u *userService) GetTasksForUser(userID string) ([]taskService.Task, error) {
	var allTasks []taskService.Task

	allTasks, err := u.repo.GetTasksForUser(userID)
	if err != nil {
		return nil, err
	}

	return allTasks, nil
}

func (u *userService) CreateUser(req UserRequest) (User, error) {
	var user User

	if req.Email == "" || req.Password == "" {
		return User{}, errors.New("fill in all the fields")
	}

	user = User{
		Id:       uuid.New().String(),
		Email:    req.Email,
		Password: req.Password,
	}

	err := u.repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *userService) UpdateUser(id string, req UserRequest) (User, error) {
	if req.Email == "" || req.Password == "" {
		return User{}, errors.New("fill in all the fields")
	}

	user, err := u.repo.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	user.Email = req.Email
	user.Password = req.Password

	err = u.repo.UpdateUser(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *userService) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}
