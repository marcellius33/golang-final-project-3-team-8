package params

import (
	"kanbanboard/models"
	"time"
)

type UserRegisterResponse struct {
	ID        uint      `json:"id"  example:"1"`
	FullName  string    `json:"full_name" example:"curry"`
	Email     string    `json:"email" example:"curry@gmail.com"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateResponse struct {
	ID        uint      `json:"id"  example:"1"`
	FullName  string    `json:"full_name" example:"curry"`
	Email     string    `json:"email" example:"curry@gmail.com"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func ParseToCreateUserResponse(user *models.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func ParseToUpdateUserResponse(user *models.User) UserUpdateResponse {
	return UserUpdateResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}
}
