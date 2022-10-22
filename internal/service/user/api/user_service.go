package userService

import (
	"belajar-go-echo/internal/dto"
	"context"
)

type UserService interface {
	GetAllUsers(ctx context.Context) (dto.Users, error)
	CreateUser(ctx context.Context, user dto.UserRequest) (dto.User, error)
}
