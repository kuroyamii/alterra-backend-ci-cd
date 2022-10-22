package authServiceMock

import (
	"belajar-go-echo/internal/dto"
	authRepository "belajar-go-echo/internal/repository/auth/api"
	authRepositoryMock "belajar-go-echo/internal/repository/auth/mock"
	"context"
)

type AuthServiceMock struct {
	ar authRepository.AuthRepository
}

func NewAuthServiceMock() AuthServiceMock {
	ar := authRepositoryMock.NewAuthRepoMock()
	return AuthServiceMock{
		ar: ar,
	}
}

func (ar AuthServiceMock) CreateJWT(ctx context.Context, request dto.UserRequest) (dto.AuthResponse, error) {
	response := dto.AuthResponse{
		Token: "token",
	}
	return response, nil
}
