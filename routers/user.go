package routers

import (
	"github.com/gin-gonic/gin"
	"kanbanboard/controllers"
)

func InitUserRoutes(Routes *gin.Engine, controller *controllers.UserController) {
	userRouter := Routes.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegisterController)
		userRouter.POST("/login", controller.UserLoginController)
	}
}
