package userRepository

import (
	"belajar-go-echo/internal/entity"
	"context"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) (entity.Users, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
}
