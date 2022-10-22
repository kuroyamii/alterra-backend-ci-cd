package authService

import (
	"belajar-go-echo/internal/dto"
	authRepository "belajar-go-echo/internal/repository/auth/api"
	"belajar-go-echo/pkg/config"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type authService struct {
	ar authRepository.AuthRepository
}

func NewAuthService(ar authRepository.AuthRepository) authService {
	return authService{
		ar: ar,
	}
}

func createJWTToken(email string, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["password"] = password
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_KEY))
}

func (ar authService) CreateJWT(ctx context.Context, request dto.UserRequest) (dto.AuthResponse, error) {
	err := ar.ar.ValidateCredentials(ctx, request.Email, request.Password)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	token, err := createJWTToken(request.Email, request.Password)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	response := dto.AuthResponse{
		Token: token,
	}
	return response, nil
}
