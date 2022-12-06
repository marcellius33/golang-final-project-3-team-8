package routers

import (
	"kanbanboard/controllers"
	"kanbanboard/middlewares"

	"github.com/gin-gonic/gin"
)

func InitCategoriesRoutes(Routes *gin.Engine, controller *controllers.CategoryController) {
	categoryRouter := Routes.Group("/categories")
	{
		categoryRouter.POST("/", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}),controller.CreateCategory)
		categoryRouter.GET("/", controller.GetCategories)
		categoryRouter.PATCH("/:categoryId", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), controller.DeleteCategory)
	}
}
