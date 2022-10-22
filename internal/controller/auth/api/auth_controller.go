package authController

import "github.com/labstack/echo/v4"

type AuthController interface {
	HandleLogin(c echo.Context) error
}
