package authRepository

import (
	"belajar-go-echo/internal/entity"
	"context"

	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authRepository {
	return authRepository{
		DB: db,
	}
}

func (ar authRepository) ValidateCredentials(ctx context.Context, email string, password string) error {
	user := entity.User{}
	err := ar.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
