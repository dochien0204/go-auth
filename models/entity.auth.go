package models

import "gorm.io/gorm"

type EntityUser struct {
	gorm.Model
	Username           string `gorm:"size:255;not null;unique" json:"username"`
	Password           string `gorm:"size:255;not null" json:"password"`
	PasswordResetToken string `gorm:"size:255; not null" json:"password_reset_token"`
	AccessToken        string `json:"accessToken"`
	Email              string `json:"email"`
}
