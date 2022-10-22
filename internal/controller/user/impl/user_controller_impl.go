package userController

import (
	"belajar-go-echo/internal/dto"
	userService "belajar-go-echo/internal/service/user/api"

	"github.com/labstack/echo/v4"
)

type userController struct {
	router *echo.Echo
	us     userService.UserService
}

func NewUserController(router *echo.Echo, us userService.UserService) userController {
	return userController{
		router: router,
		us:     us,
	}
}

func (uc userController) HandleGetAllUsers(c echo.Context) error {
	res, err := uc.us.GetAllUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": res,
	})
}

func (uc userController) HandleCreateUser(c echo.Context) error {
	userDto := dto.UserRequest{}
	err := c.Bind(&userDto)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"error": err.Error(),
		})
	}
	res, err := uc.us.CreateUser(c.Request().Context(), userDto)
	if err != nil {
		return echo.NewHTTPError(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": res,
	})
}
