package main

import (
	"flag"
	"kanbanboard/controllers"
	"kanbanboard/database"
	"kanbanboard/database/seed"
	_ "kanbanboard/initializer"
	"kanbanboard/routers"
	"kanbanboard/services"
	"os"

	"github.com/gin-gonic/gin"

	"kanbanboard/repositories"
)

// TODO: Middleware

func init() {
	database.Connect()
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			userRepository := repositories.NewUserRepository(database.GetDB())
			userSeed := seed.NewUserSeeder(userRepository)
			userSeed.Execute()
		}
	}
}

func main() {
	handleArgs()

	Routes := gin.Default()

	userRepository := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	routers.InitUserRoutes(Routes, userController)

	categoryRepository := repositories.NewCategoryRepository(database.GetDB())
	categoryService := services.NewCategoryService(categoryRepository, userRepository)
	categoryController := controllers.NewCategoryController(categoryService)
	routers.InitCategoriesRoutes(Routes, categoryController)

	Routes.Run(os.Getenv("SERVER_PORT"))
}