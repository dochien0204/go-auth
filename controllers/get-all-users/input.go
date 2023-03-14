package getallusers

type InputUpdateUser struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	PasswordResetToken string `json:"passwordResetToken"`
	Email              string `json:"email"`
}
