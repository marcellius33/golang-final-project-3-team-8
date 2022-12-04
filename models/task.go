package models

import (
	"time"
)

// Title: not null
// Description: not null
// Status: true or false (boolean)

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey;type:integer"`
	Title       string `json:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	Status      bool   `json:"status" gorm:"type:boolean"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
	User        User
	Category    Category
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
