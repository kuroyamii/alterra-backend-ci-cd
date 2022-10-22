package userService

import (
	"belajar-go-echo/internal/dto"
	"belajar-go-echo/internal/entity"
	userRepository "belajar-go-echo/internal/repository/user/api"
	"context"

	"gorm.io/gorm"
)

type userService struct {
	ur userRepository.UserRepository
}

func NewUserService(ur userRepository.UserRepository) userService {
	return userService{
		ur: ur,
	}
}

func (us userService) GetAllUsers(ctx context.Context) (dto.Users, error) {
	users, err := us.ur.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	usersResponse := dto.Users{}
	for _, user := range users {
		singleUser := dto.User{
			Model: &gorm.Model{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
			Email:    user.Email,
			Password: user.Password,
		}
		usersResponse = append(usersResponse, singleUser)
	}
	return usersResponse, nil
}
func (us userService) CreateUser(ctx context.Context, user dto.UserRequest) (dto.User, error) {
	userEntity := entity.User{
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := us.ur.CreateUser(ctx, userEntity)
	if err != nil {
		return dto.User{}, nil
	}

	userDto := dto.User{
		Model: &gorm.Model{
			ID:        res.ID,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
			DeletedAt: res.DeletedAt,
		},
		Email:    res.Email,
		Password: res.Password,
	}
	return userDto, nil
}
