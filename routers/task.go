package routers

import (
	"kanbanboard/controllers"
	"kanbanboard/middlewares"

	"github.com/gin-gonic/gin"
)

func InitTaskRoutes(Routes *gin.Engine, controller *controllers.TaskController) {
	taskRouter := Routes.Group("/tasks")
	{
		taskRouter.POST("/", middlewares.Authentication(), middlewares.Authorization([]string{"admin", "member"}),controller.CreateTask)
		taskRouter.GET("/", middlewares.Authentication(), middlewares.Authorization([]string{"admin", "member"}),controller.GetTasks)
		taskRouter.PUT("/:taskId", middlewares.Authentication(), middlewares.Authorization([]string{"admin", "member"}), controller.UpdateTask)
		taskRouter.PATCH("/update-status/:taskId", middlewares.Authentication(), middlewares.Authorization([]string{"admin", "member"}), controller.UpdateTaskStatus)
		taskRouter.PATCH("/update-category/:taskId", middlewares.Authentication(), middlewares.Authorization([]string{"admin", "member"}), controller.UpdateTaskCategory)
		taskRouter.DELETE("/:taskId", middlewares.Authentication(), middlewares.Authorization([]string{"admin", "member"}), controller.DeleteTask)
	}
}
