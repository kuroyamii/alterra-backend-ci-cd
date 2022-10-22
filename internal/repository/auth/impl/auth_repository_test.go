package authRepository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error %v was encountered while opening a mock database connection\n", err.Error())
	}
	dbGorm, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))
	return dbGorm, mock, err
}

func iniMockRepo(dbGorm *gorm.DB) (authRepository, *echo.Echo) {
	router := echo.New()
	var authRepository = authRepository{
		DB: dbGorm,
	}
	return authRepository, router
}

func TestValidateCreds(t *testing.T) {
	dbGorm, mock, err := initMockDB(t)
	if err != nil {
		t.Errorf("Error Opening DB Connection: %v", err.Error())
	}
	repo, _ := iniMockRepo(dbGorm)

	testCases := []struct {
		Name     string
		Email    string
		Password string
		Context  context.Context
	}{
		{
			Name:     "test case 1",
			Context:  context.TODO(),
			Email:    "email1",
			Password: "password1",
		},
		{
			Name:     "test case 2",
			Context:  context.TODO(),
			Email:    "email2",
			Password: "password2",
		},
	}

	row := mock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "email", "password"}).AddRow(1, time.Now(), time.Now(), nil, "email1", "password1").AddRow(2, time.Now(), time.Now(), nil, "email2", "password2")
	for _, val := range testCases {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (email = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).WithArgs(val.Email, val.Password).WillReturnRows(row)
		err := repo.ValidateCredentials(val.Context, val.Email, val.Password)
		t.Log(err)
	}
}

func TestNewAuthRepo(t *testing.T) {
	dbGorm, _, err := initMockDB(t)
	assert.NoError(t, err)
	repo := NewAuthRepository(dbGorm)
	assert.NotEmpty(t, repo)
}
