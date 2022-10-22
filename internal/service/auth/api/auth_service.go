package authService

import (
	"belajar-go-echo/internal/dto"
	"context"
)

type AuthService interface {
	CreateJWT(ctx context.Context, request dto.UserRequest) (dto.AuthResponse, error)
}
