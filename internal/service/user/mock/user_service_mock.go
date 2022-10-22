package userServiceMock

import (
	"belajar-go-echo/internal/dto"
	userRepository "belajar-go-echo/internal/repository/user/api"
	UserRepositoryMock "belajar-go-echo/internal/repository/user/mock"
	"context"
)

type UserServiceMock struct {
	ur userRepository.UserRepository
}

func NewUserServiceMock() UserServiceMock {
	mockRepo := UserRepositoryMock.NewUserRepoMock()
	return UserServiceMock{
		ur: mockRepo,
	}
}

func (us UserServiceMock) GetAllUsers(ctx context.Context) (dto.Users, error) {
	users := dto.Users{}
	user1 := dto.User{
		Email:    "email1",
		Password: "password1",
	}
	user2 := dto.User{
		Email:    "email1",
		Password: "password1",
	}
	users = append(users, user1)
	users = append(users, user2)
	return users, nil
}
func (us UserServiceMock) CreateUser(ctx context.Context, user dto.UserRequest) (dto.User, error) {
	res := dto.User{
		Email:    user.Email,
		Password: user.Password,
	}
	return res, nil
}
