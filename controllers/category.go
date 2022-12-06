package controllers

import (
	"kanbanboard/helpers"
	"kanbanboard/params"
	"kanbanboard/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (p *CategoryController) CreateCategory(c *gin.Context) {
	CategoryRequest := params.CreateCategoryRequest{}
	if err := c.ShouldBindJSON(&CategoryRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	createCategory, err := p.service.CreateCategory(CategoryRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessCreateResponse(createCategory, "Create Category Success"))
}

func (p *CategoryController) GetCategories(c *gin.Context) {
	Categories, err := p.service.GetCategories()
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(Categories, "Get Categories Success"))
}

func (p *CategoryController) UpdateCategory(c *gin.Context) {
	categoryRequest := params.UpdateCategoryRequest{}
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	categoryId, err := strconv.Atoi(c.Param("categoryId"))

	updateCategory, err := p.service.UpdateCategory(uint(categoryId), categoryRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(updateCategory, "Update Category Success"))
}

func (p *CategoryController) DeleteCategory(c *gin.Context) {
	CategoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	err = p.service.DeleteCategory(uint(CategoryId))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.DeleteSuccess("Your Category has been successfully deleted"))
}
