package forgot

import (
	"fmt"
	email "jwt-project/controllers/email"
	otptokens "jwt-project/controllers/otp_tokens"
	"jwt-project/models"
	"jwt-project/utils"
)

type service struct {
	repository Repository
	otptokens.IOtp
}

func NewForgotService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ForgotService(input *InputForgot) (string, bool) {
	var user models.EntityUser
	user.Username = input.Username

	resultForgot, errForgot := s.repository.ForgotRepository(&user)

	if !errForgot {
		return "Account is error", false
	}

	otpCode, err := utils.GenCaptchaCode()

	s.GenerateNewOTP(otpCode)
	if err != nil {
		return "Send OTP failed", false
	}

	fmt.Println(otpCode)

	//start mail
	mail := email.NewMail()
	message := "OTP: " + otpCode

	//sent mail
	mail.SendSimpleMessage(resultForgot.Email, message)

	return "OTP is sent in your mail. Please check that", true
}
