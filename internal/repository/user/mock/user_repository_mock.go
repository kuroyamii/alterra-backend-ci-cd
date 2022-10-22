package UserRepositoryMock

import (
	"belajar-go-echo/internal/entity"
	"context"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepositoryMock struct {
	DB *gorm.DB
}

func initMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, _ := sqlmock.New()
	dbGorm, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))
	return dbGorm, mock, err
}

func NewUserRepoMock() UserRepositoryMock {
	db, _, _ := initMockDB()
	return UserRepositoryMock{
		DB: db,
	}
}

func (ur UserRepositoryMock) GetAllUsers(ctx context.Context) (entity.Users, error) {
	users := entity.Users{}
	user1 := entity.User{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     "email1",
		Password:  "password1",
	}
	user2 := entity.User{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     "email1",
		Password:  "password1",
	}
	users = append(users, user1)
	users = append(users, user2)
	return users, nil
}
func (ur UserRepositoryMock) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	res := entity.User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     user.Email,
		Password:  user.Password,
	}
	return res, nil
}
