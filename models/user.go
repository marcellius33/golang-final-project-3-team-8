package models

import (
	"time"
)

// Email: format valid, unique index, not null
// FullName: not null
// Password: not null, minimum length 6
// Role: not null, only admin or member

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;type:integer"`
	FullName  string    `json:"full_name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Role      string    `json:"role" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
