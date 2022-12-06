package params

import (
	"kanbanboard/models"
	"time"
)

type CreateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

type UpdateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

type CreateCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

func ParseToCreateCategoryResponse(category *models.Category) CreateCategoryResponse {
	return CreateCategoryResponse{
		ID:        category.ID,
		Type:	   category.Type,
		CreatedAt: category.CreatedAt,
	}
}

type UpdateCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

func ParseToUpdateCategoryResponse(category *models.Category) UpdateCategoryResponse {
	return UpdateCategoryResponse{
		ID:        category.ID,
		Type:	   category.Type,
		CreatedAt: category.CreatedAt,
	}
}