package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `gorm:"size: 255;not null; unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Email    string `gorm:"size:255;" json:"email"`
}
