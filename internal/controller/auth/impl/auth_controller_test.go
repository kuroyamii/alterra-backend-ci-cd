package authController

import (
	authServiceMock "belajar-go-echo/internal/service/auth/mock"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func initMockController(mockService authServiceMock.AuthServiceMock) authController {
	router := echo.New()
	var authController = authController{
		router: router,
		as:     mockService,
	}
	return authController
}

func TestHandleLogin(t *testing.T) {
	service := authServiceMock.NewAuthServiceMock()
	controller := initMockController(service)

	testCases := []struct {
		Name     string
		Method   string
		Body     map[string]interface{}
		HasError bool
	}{
		{
			Name:   "test case 1",
			Method: "POST",
			Body: map[string]interface{}{
				"email":    "email1",
				"password": "password1",
			},
			HasError: false,
		},
		{
			Name:   "test case 2",
			Method: "POST",
			Body: map[string]interface{}{
				"email":    "email2",
				"password": 497293784,
			},
			HasError: true,
		},
	}

	for _, val := range testCases {
		t.Run(val.Name, func(t *testing.T) {
			res, _ := json.Marshal(val.Body)
			r := httptest.NewRequest(val.Method, "/login", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			r.Header.Set("Content-Type", "application/json")
			ctx := controller.router.NewContext(r, w)
			err := controller.HandleLogin(ctx)
			if val.HasError == false {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestNewAuthController(t *testing.T) {
	router := echo.New()
	service := authServiceMock.NewAuthServiceMock()
	controller := NewAuthController(service, router)
	assert.NotEmpty(t, controller)
}
