package otptokens

import "jwt-project/models"

type service struct {
	repository IOTPRepository
}

func NewService(repository IOTPRepository) *service {
	return &service{
		repository: repository,
	}
}
func (s *service) GenerateNewOTP(token string, user uint) *models.OTPToken {
	return s.repository.GenerateNewOTP(token, user)
}
