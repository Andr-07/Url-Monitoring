package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   uint
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
