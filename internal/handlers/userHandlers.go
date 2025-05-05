package handlers

import (
	"context"
	"firstCoursePractice/internal/userService"
	"firstCoursePractice/internal/web/users"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *userHandler {
	return &userHandler{service: service}
}

func (u *userHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, recievedUser := range allUsers {
		user := users.User{
			Id:    &recievedUser.Id,
			Email: &recievedUser.Email,
		}

		response = append(response, user)
	}

	return response, nil
}

func (u *userHandler) GetUsersUserIdTasks(_ context.Context, request users.GetUsersUserIdTasksRequestObject) (users.GetUsersUserIdTasksResponseObject, error) {
	userId := request.UserId
	allTasks, err := u.service.GetTasksForUser(userId)
	if err != nil {
		return nil, err
	}

	response := users.GetUsersUserIdTasks200JSONResponse{}

	for _, recievedTask := range allTasks {
		task := users.Task{
			Id:     &recievedTask.ID,
			Task:   &recievedTask.Task,
			IsDone: &recievedTask.IsDone,
			UserId: &userId,
		}
		response = append(response, task)
	}

	return response, nil
}

func (u *userHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	req := request.Body
	if req.Email == nil || req.Password == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "email and password are required fields")
	}

	userToCreate := userService.UserRequest{
		Email:    *req.Email,
		Password: *req.Password,
	}

	user, err := u.service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:    &user.Id,
		Email: &user.Email,
	}

	return response, nil
}

func (u *userHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	err := u.service.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	response := users.DeleteUsersId204Response{}

	return response, nil
}

func (u *userHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id

	req := request.Body
	if req.Email == nil || req.Password == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "email and password are required fields")
	}

	userToUpdate := userService.UserRequest{
		Email:    *req.Email,
		Password: *req.Password,
	}

	user, err := u.service.UpdateUser(id, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:    &user.Id,
		Email: &user.Email,
	}

	return response, nil
}
