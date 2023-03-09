package forgot

import (
	"jwt-project/models"
)

type service struct {
	repository Repository
}

func NewForgotService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ForgotService(input *InputForgot) (*models.EntityUser, string) {
	var user models.EntityUser
	user.Username = input.Username
	user.PasswordResetToken = input.ForgotToken

	resultForgot, errForgot := s.repository.ForgotRepository(&user)

	return resultForgot, errForgot
}
