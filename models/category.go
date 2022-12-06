package models

import (
	"time"
)

// Type: not null

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey;type:integer"`
	Type      string    `json:"type" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Task	  []Task
}
