package models

import (
	"time"
)

// Title: not null
// Description: not null
// Status: true or false (boolean)

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey;type:integer"`
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	Status      bool   `json:"status" gorm:"type:boolean" default:"false"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
