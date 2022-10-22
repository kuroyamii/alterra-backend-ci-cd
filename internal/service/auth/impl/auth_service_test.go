package authService

import (
	"belajar-go-echo/internal/dto"
	authRepositoryMock "belajar-go-echo/internal/repository/auth/mock"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initMockService(mockRepo authRepositoryMock.AuthRepoMock) authService {

	var authService = authService{
		ar: mockRepo,
	}
	return authService
}

func TestCreateJWT(t *testing.T) {
	repo := authRepositoryMock.NewAuthRepoMock()
	service := initMockService(repo)

	testCases := []struct {
		Name     string
		Context  context.Context
		Email    string
		Password string
		Token    string
		HasError bool
	}{
		{
			Name:     "test case 1",
			Context:  context.TODO(),
			Token:    "token",
			HasError: false,
			Email:    "email1",
			Password: "password1",
		},
		{
			Name:     "test case 2",
			Context:  context.TODO(),
			Token:    "token",
			HasError: true,
			Email:    "slkajflkajsf",
			Password: "alskjdflksjdf",
		},
	}
	for _, val := range testCases {
		req := dto.UserRequest{
			Email:    val.Email,
			Password: val.Password,
		}
		_, err := service.CreateJWT(val.Context, req)
		if val.HasError == true {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestCreateToken(t *testing.T) {
	testCases := []struct {
		Name     string
		Context  context.Context
		Email    string
		Password string
		Token    string
		HasError bool
	}{
		{
			Name:     "test case 1",
			Context:  context.TODO(),
			Token:    "token",
			HasError: false,
			Email:    "email1",
			Password: "password1",
		},
		{
			Name:     "test case 2",
			Context:  context.TODO(),
			Token:    "token",
			HasError: true,
			Email:    "slkajflkajsf",
			Password: "alskjdflksjdf",
		},
	}

	for _, val := range testCases {
		t.Run(val.Name, func(t *testing.T) {
			token, err := createJWTToken(val.Email, val.Password)
			assert.NoError(t, err)
			assert.NotEmpty(t, token)
		})
	}
}

func TestNewAuthService(t *testing.T) {
	repo := authRepositoryMock.NewAuthRepoMock()
	service := NewAuthService(repo)
	assert.NotEmpty(t, service)
}
