package main

import (
	authController "belajar-go-echo/internal/controller/auth/impl"
	userController "belajar-go-echo/internal/controller/user/impl"
	"belajar-go-echo/internal/entity"
	authRepository "belajar-go-echo/internal/repository/auth/impl"
	userRepository "belajar-go-echo/internal/repository/user/impl"
	routes "belajar-go-echo/internal/routes"
	authService "belajar-go-echo/internal/service/auth/impl"
	userService "belajar-go-echo/internal/service/user/impl"
	"belajar-go-echo/pkg/config"
	"belajar-go-echo/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	env := config.GetEnvVariables()

	router := config.InitRouter()
	db := config.InitDatabase(env)
	config.InitialMigration(db, &entity.User{})

	userRepository := userRepository.NewRepository(db)
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(router, userService)
	userRoutes := routes.NewUserRoutes(userController, router)
	userRoutes.InitEndpoints()

	authRepository := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepository)
	authController := authController.NewAuthController(authService, router)
	authRoutes := routes.NewAuthRoutes(authController, router)
	authRoutes.InitEndpoints()

	server := server.NewServer(env["SERVER_ADDRESS"], router)
	server.ListenAndServe()
}
