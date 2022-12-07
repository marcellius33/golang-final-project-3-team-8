package services

import (
	"kanbanboard/models"
	"kanbanboard/params"
	"kanbanboard/repositories"
)

type CategoryService interface {
	CreateCategory(createCategoryRequest params.CreateCategoryRequest) (*params.CreateCategoryResponse, error)
	GetCategories() (*[]models.Category, error)
	UpdateCategory(categoryId uint, categoryUpdateRequest params.UpdateCategoryRequest) (*params.UpdateCategoryResponse, error)
	DeleteCategory(categoryId uint) error
}

type categoryService struct {
	categoryR repositories.CategoryRepository
	userR    repositories.UserRepository
}

func NewCategoryService(categoryR repositories.CategoryRepository, userR repositories.UserRepository) CategoryService {
	return &categoryService{
		categoryR: categoryR,
		userR:    userR,
	}
}

func (c *categoryService) CreateCategory(createCategoryRequest params.CreateCategoryRequest) (*params.CreateCategoryResponse, error) {
	newCategory := models.Category{
		Type: createCategoryRequest.Type,
	}

	_, err := c.categoryR.CreateCategory(&newCategory)

	if err != nil {
		return &params.CreateCategoryResponse{}, err
	}
	resp := params.ParseToCreateCategoryResponse(&newCategory)

	return &resp, nil
}

func (c *categoryService) GetCategories() (*[]models.Category, error) {
	var categories []models.Category
	_, err := c.categoryR.GetCategories(&categories)

	if err != nil {
		return &[]models.Category{}, err
	}

	return &categories, nil
}

func (c *categoryService) UpdateCategory(categoryId uint, categoryUpdateRequest params.UpdateCategoryRequest) (*params.UpdateCategoryResponse, error) {
	CategoryModel, err := c.categoryR.FindCategoryById(categoryId)
	if err != nil {
		return &params.UpdateCategoryResponse{}, err
	}
	CategoryModel.Type = categoryUpdateRequest.Type

	_, err = c.categoryR.UpdateCategory(categoryId, CategoryModel)

	if err != nil {
		return &params.UpdateCategoryResponse{}, err
	}
	resp := params.ParseToUpdateCategoryResponse(CategoryModel)

	return &resp, nil
}

func (c *categoryService) DeleteCategory(categoryId uint) error {
	return c.categoryR.DeleteCategory(categoryId)
}

