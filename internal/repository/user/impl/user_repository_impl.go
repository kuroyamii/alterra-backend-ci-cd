package userRepository

import (
	"belajar-go-echo/internal/entity"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) userRepository {
	return userRepository{
		DB: db,
	}
}

func (ur userRepository) GetAllUsers(ctx context.Context) (entity.Users, error) {
	users := make(entity.Users, 0)
	err := ur.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (ur userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := ur.DB.Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
