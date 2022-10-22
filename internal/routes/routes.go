package routes

import (
	authController "belajar-go-echo/internal/controller/auth/api"
	userController "belajar-go-echo/internal/controller/user/api"
	"belajar-go-echo/pkg/middleware"

	"github.com/labstack/echo/v4"
)

type UserRoutes struct {
	uc     userController.UserController
	router *echo.Echo
}

type AuthRoutes struct {
	ac     authController.AuthController
	router *echo.Echo
}

func NewAuthRoutes(ac authController.AuthController, router *echo.Echo) AuthRoutes {
	return AuthRoutes{
		ac:     ac,
		router: router,
	}
}

func NewUserRoutes(uc userController.UserController, router *echo.Echo) UserRoutes {
	return UserRoutes{
		uc:     uc,
		router: router,
	}
}

func (urs UserRoutes) InitEndpoints() {
	userRouters := urs.router.Group("/users")
	userRouters.Use(middleware.ValidateJWT)
	userRouters.GET("", urs.uc.HandleGetAllUsers)

	urs.router.POST("/users", urs.uc.HandleCreateUser)
}

func (ars AuthRoutes) InitEndpoints() {
	ars.router.POST("/login", ars.ac.HandleLogin)
}
