package models

import "time"

type OTPToken struct {
	ID        int       `gorm:"json:otpId"`
	Token     string    `gorm:"size:255;not null" json:"otpToken"`
	ExpiredAt time.Time `json:"expiredAt"`
}
