package routers

import (
	"github.com/gin-gonic/gin"
	"kanbanboard/controllers"
	"kanbanboard/middlewares"
)

func InitUserRoutes(Routes *gin.Engine, controller *controllers.UserController) {
	userRouter := Routes.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegisterController)
		userRouter.POST("/login", controller.UserLoginController)
		userRouter.PUT("/update-account", middlewares.Authentication(), controller.UserUpdateController)
		userRouter.DELETE("/delete-account", middlewares.Authentication(), controller.UserDeleteController)
	}
}
