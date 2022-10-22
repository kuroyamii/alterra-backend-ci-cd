package userService

import (
	"belajar-go-echo/internal/dto"
	userRepositoryMock "belajar-go-echo/internal/repository/user/mock"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initMockService(mockRepo userRepositoryMock.UserRepositoryMock) userService {

	var userService = userService{
		ur: mockRepo,
	}
	return userService
}

func TestGetUsers(t *testing.T) {
	repo := userRepositoryMock.NewUserRepoMock()
	service := initMockService(repo)

	testCases := []struct {
		Name    string
		Context context.Context
		Users   dto.Users
	}{
		{
			Name:    "test case 1",
			Context: context.TODO(),
			Users: dto.Users{
				dto.User{
					Email:    "email1",
					Password: "password1",
				},
				dto.User{
					Email:    "email2",
					Password: "password2",
				},
			},
		},
	}
	for _, val := range testCases {
		t.Run(val.Name, func(t *testing.T) {
			users, err := service.GetAllUsers(val.Context)
			assert.NoError(t, err)
			assert.NotEmpty(t, users)

		})
	}
}

func TestCreateUser(t *testing.T) {
	repo := userRepositoryMock.NewUserRepoMock()
	service := initMockService(repo)

	testCases := []struct {
		Name    string
		Context context.Context

		Email    string
		Password string
	}{
		{
			Name:     "test case 1",
			Context:  context.TODO(),
			Email:    "email1",
			Password: "password1",
		},
		{
			Name:     "test case 2",
			Context:  context.TODO(),
			Email:    "email2",
			Password: "password2",
		},
	}

	for _, val := range testCases {
		t.Run(val.Name, func(t *testing.T) {
			userReq := dto.UserRequest{
				Email:    val.Email,
				Password: val.Password,
			}
			res, err := service.CreateUser(val.Context, userReq)
			assert.NoError(t, err)
			assert.Equal(t, res.Email, val.Email)
			assert.Equal(t, res.Password, val.Password)
		})
	}
}

func TestNewUserService(t *testing.T) {
	repo := userRepositoryMock.NewUserRepoMock()
	service := NewUserService(repo)
	assert.NotEmpty(t, service)
}
