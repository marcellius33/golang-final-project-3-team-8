package main

import (
	"github.com/gin-gonic/gin"
	"kanbanboard/controllers"
	"kanbanboard/database"
	_ "kanbanboard/initializer"
	"kanbanboard/routers"
	"kanbanboard/services"
	"os"

	"kanbanboard/repositories"
)

// TODO: Seed Admin
// TODO: Check Validation
// TODO: Update and Delete User
// TODO: Middleware

func init() {
	database.Connect()
}

func main() {
	Routes := gin.Default()

	userRepository := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	routers.InitUserRoutes(Routes, userController)

	Routes.Run(os.Getenv("SERVER_PORT"))
}
