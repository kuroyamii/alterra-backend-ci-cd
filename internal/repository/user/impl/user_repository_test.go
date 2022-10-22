package userRepository

import (
	"belajar-go-echo/internal/entity"
	"context"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

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

func iniMockRepo(dbGorm *gorm.DB) (userRepository, *echo.Echo) {
	router := echo.New()
	var userRepository = userRepository{
		DB: dbGorm,
	}
	return userRepository, router
}

func TestGetUsers(t *testing.T) {
	dbGorm, mock, err := initMockDB(t)
	if err != nil {
		t.Errorf("Error Opening DB Connection: %v", err.Error())
	}
	repo, _ := iniMockRepo(dbGorm)

	testCases := []struct {
		Name     string
		Expected []map[string]interface{}
		Context  context.Context
	}{
		{
			Name:    "test case 1",
			Context: context.TODO(),
			Expected: []map[string]interface{}{
				{
					"email":    "email1",
					"passwrod": "password1",
				},
				{
					"email":    "email2",
					"password": "password2",
				},
			},
		},
	}

	row := mock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "email", "password"}).AddRow(1, time.Now(), time.Now(), nil, "email1", "password1").AddRow(2, time.Now(), time.Now(), nil, "email2", "password2")
	for _, val := range testCases {
		t.Run(val.Name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).WillReturnRows(row)
			users, err := repo.GetAllUsers(val.Context)
			assert.NoError(t, err)
			assert.NotEmpty(t, users)
		})
	}
}

func TestCreateUser(t *testing.T) {
	dbGorm, mock, err := initMockDB(t)
	if err != nil {
		t.Errorf("Error Opening DB Connection: %v", err.Error())
	}
	repo, _ := iniMockRepo(dbGorm)

	testCases := []struct {
		Name     string
		Email    string
		Password string
		ID       int64
		Expected map[string]interface{}
		Context  context.Context
	}{
		{
			Name:     "test case 1",
			Context:  context.TODO(),
			Email:    "email1",
			Password: "password1",
			ID:       1,
			Expected: map[string]interface{}{

				"email":    "email1",
				"password": "password1",
			},
		},
		{
			Name:     "test case 2",
			Context:  context.TODO(),
			Email:    "email2",
			Password: "password2",
			ID:       2,
			Expected: map[string]interface{}{

				"email":    "email2",
				"password": "password2",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.Name, func(t *testing.T) {
			mock.ExpectBegin()
			result := sqlmock.NewResult(val.ID, 0)
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`) VALUES (?,?,?,?,?)")).WithArgs(AnyTime{}, AnyTime{}, nil, val.Email, val.Password).WillReturnResult(result)
			mock.ExpectCommit()
			req := entity.User{
				Email:    val.Email,
				Password: val.Password,
			}

			user, err := repo.CreateUser(val.Context, req)
			assert.NoError(t, err)
			assert.Equal(t, user.Email, val.Email)
			assert.Equal(t, user.Password, val.Password)
		})
	}
}

func TestNewRepo(t *testing.T) {
	dbGorm, _, err := initMockDB(t)
	assert.NoError(t, err)
	repo := NewRepository(dbGorm)
	assert.NotEmpty(t, repo)
}
