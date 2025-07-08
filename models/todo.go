package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	Title     string `gorm:"not null"`
	Done      bool   `gorm:"default:false"`
	IsDeleted bool   `gorm:"default:false"`
	gorm.Model
	UserID uint `gorm:"not null"`
}
