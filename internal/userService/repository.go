package userService

import (
	"firstCoursePractice/internal/taskService"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]User, error)
	GetUserById(id string) (User, error)
	GetTasksForUser(userID string) ([]taskService.Task, error)
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *userRepository) GetUserById(id string) (User, error) {
	var user User
	err := u.db.First(&user, "id = ?", id).Error
	return user, err
}

func (u *userRepository) GetTasksForUser(userID string) ([]taskService.Task, error) {
	var tasks []taskService.Task
	err := u.db.Find(&tasks, "user_id = ?", userID).Error
	return tasks, err
}

func (u *userRepository) CreateUser(user User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) UpdateUser(user User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) DeleteUser(id string) error {
	return u.db.Delete(&User{}, "id = ?", id).Error
}
