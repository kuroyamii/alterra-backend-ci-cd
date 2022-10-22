package AuthRepoMockMock

import (
	"context"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AuthRepoMock struct {
	DB *gorm.DB
}

func (ar AuthRepoMock) ValidateCredentials(ctx context.Context, email string, password string) error {
	if email != "email1" || password != "password1" {
		return errors.New("Error")
	}
	return nil
}

func initMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, _ := sqlmock.New()
	dbGorm, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))
	return dbGorm, mock, err
}

func NewAuthRepoMock() AuthRepoMock {
	db, _, _ := initMockDB()
	return AuthRepoMock{
		DB: db,
	}
}
