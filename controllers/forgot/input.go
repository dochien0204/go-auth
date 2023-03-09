package forgot

type InputForgot struct {
	Username    string `json:"username"`
	ForgotToken string `json:"forgotToken"`
}
