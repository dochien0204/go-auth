package presenters

type ForgotPasswordResponse struct {
	PasswordResetLink string `json:"passwordResetLink"`
}
