package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name        string `json:"name" form:"name" binding:"required" gorm:"not null"`
	Price       uint   `json:"price" form:"price" binding:"required|numeric" gorm:"not null"`
	Description string `json:"description" form:"description" binding:"required" gorm:"not null"`
	UserID      uint
	User        User `json:"-"`
}
