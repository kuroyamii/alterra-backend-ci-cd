package userController

import "github.com/labstack/echo/v4"

type UserController interface {
	HandleGetAllUsers(c echo.Context) error
	HandleCreateUser(c echo.Context) error
}
