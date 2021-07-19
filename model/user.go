package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" binding:"required" gorm:"not null"`
	Username string `json:"username" form:"username" binding:"required" gorm:"not null"`
	Password string `json:"password" form:"password" binding:"required" gorm:"not null"`
	Token    string `json:"token"`
}
