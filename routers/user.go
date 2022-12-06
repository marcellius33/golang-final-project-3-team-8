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
		userRouter.PUT("/update-account", middlewares.Authentication(), middlewares.Authorization([]string{"admin", "member"}), controller.UserUpdateController)
		userRouter.DELETE("/delete-account", middlewares.Authentication(), middlewares.Authorization([]string{"admin", "member"}), controller.UserDeleteController)
	}
}
