package userService

import (
	"firstCoursePractice/internal/taskService"
)

type User struct {
	Id       string             `json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
