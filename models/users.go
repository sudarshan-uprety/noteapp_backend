package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Full_name string `json:"full_name"`
	Email     string `json:"email" binding:"required" gorm:"unique"`
	Password  string `json:"password" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
