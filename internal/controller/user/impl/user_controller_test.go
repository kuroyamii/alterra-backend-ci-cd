package userController

import (
	"belajar-go-echo/internal/dto"
	userServiceMock "belajar-go-echo/internal/service/user/mock"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func initMockController(mockService userServiceMock.UserServiceMock) userController {
	router := echo.New()
	var userController = userController{
		us:     mockService,
		router: router,
	}
	return userController
}

func TestHandleUsers(t *testing.T) {
	service := userServiceMock.NewUserServiceMock()
	controller := initMockController(service)

	testCases := []struct {
		Name   string
		Method string
	}{
		{
			Name:   "test case 1",
			Method: "GET",
		},
	}

	for _, val := range testCases {
		r := httptest.NewRequest(val.Method, "/users", nil)
		w := httptest.NewRecorder()
		ctx := controller.router.NewContext(r, w)
		err := controller.HandleGetAllUsers(ctx)
		assert.NoError(t, err)
	}
}

func TestHandleCreateUser(t *testing.T) {
	service := userServiceMock.NewUserServiceMock()
	controller := initMockController(service)
	testCases := []struct {
		Name        string
		Method      string
		RequestBody dto.UserRequest
		Body        map[string]interface{}
		HasError    bool
	}{
		{
			Name:   "test case 1",
			Method: "POST",
			RequestBody: dto.UserRequest{
				Email:    "email1",
				Password: "password1",
			},
			Body: map[string]interface{}{
				"email":    "email1",
				"password": "pass1",
			},
			HasError: false,
		},
		{
			Name:   "test case 2",
			Method: "POST",
			RequestBody: dto.UserRequest{
				Email:    "email2",
				Password: "password2",
			},
			Body: map[string]interface{}{
				"email":    "email2",
				"password": 2234,
			},
			HasError: true,
		},
	}

	for _, val := range testCases {
		t.Run(val.Name, func(t *testing.T) {
			res, _ := json.Marshal(val.Body)
			r := httptest.NewRequest(val.Method, "/users", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			r.Header.Set("Content-Type", "application/json")
			ctx := controller.router.NewContext(r, w)
			err := controller.HandleCreateUser(ctx)
			if val.HasError == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewUserController(t *testing.T) {
	router := echo.New()
	service := userServiceMock.NewUserServiceMock()
	controller := NewUserController(router, service)
	assert.NotEmpty(t, controller)
}
