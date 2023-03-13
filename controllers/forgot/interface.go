package forgot

type IForgotService interface {
	ForgotService(input *InputForgot) (string, bool)
}
