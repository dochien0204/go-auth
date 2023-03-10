package models

type EntityUser struct {
	Username           string `gorm:"size:255;not null;unique" json:"username"`
	Password           string `gorm:"size:255;not null" json:"password"`
	PasswordResetToken string `gorm:"size:255; not null" json:"passwordResetToken"`
	AccessToken        string `json:"accessToken"`
}
