package authController

import (
	"belajar-go-echo/internal/dto"
	authService "belajar-go-echo/internal/service/auth/api"

	"github.com/labstack/echo/v4"
)

type authController struct {
	as     authService.AuthService
	router *echo.Echo
}

func NewAuthController(as authService.AuthService, router *echo.Echo) authController {
	return authController{
		as:     as,
		router: router,
	}
}

func (ac authController) HandleLogin(c echo.Context) error {
	req := dto.UserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"error": err.Error(),
		})
	}
	res, err := ac.as.CreateJWT(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": res,
	})
}
