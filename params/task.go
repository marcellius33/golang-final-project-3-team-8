package params

import (
	"kanbanboard/models"
	"time"
)

type CreateTaskRequest struct {
	Title		string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID	uint   `json:"category_id" binding:"required"`
}

type UpdateTaskRequest struct {
	Title		string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTaskStatusRequest struct {
	Status bool `json:"status" binding:"required"`
}

type UpdateTaskCategoryRequest struct {
	CategoryID	uint   `json:"category_id" binding:"required"`
}

type CreateTaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" `
	Description string `json:"description"`
	Status      bool   `json:"status"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func ParseToCreateTaskResponse(task *models.Task) CreateTaskResponse {
	return CreateTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		CreatedAt:   task.CreatedAt,
	}
}

type UpdateTaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" `
	Description string `json:"description"`
	Status      bool   `json:"status"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func ParseToUpdateTaskResponse(task *models.Task) UpdateTaskResponse {
	return UpdateTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		CreatedAt:   task.CreatedAt,
	}
}

type GetTaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" `
	Description string `json:"description"`
	Status      bool   `json:"status"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	User      struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
	}
}

func ParseToGetTaskResponse(task models.Task, user models.User) GetTaskResponse {
	return GetTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		CreatedAt:   task.CreatedAt,
		User: struct {
			ID       uint   `json:"id"`
			Email    string `json:"email"`
			FullName string `json:"full_name"`
		}{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
		},
	}
}

func ParseToGetTasksResponse(tasks []models.Task, user models.User) []GetTaskResponse {
	var responses []GetTaskResponse
	for _, task := range tasks {
		responses = append(responses, ParseToGetTaskResponse(task, user))
	}
	return responses
}