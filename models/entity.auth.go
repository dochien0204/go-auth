package models

type EntityUser struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	PasswordResetToken string `json:"passwordResetToken"`
}
