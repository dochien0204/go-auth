package register

import "jwt-project/models"

type service struct {
	repository Repository
}

func NewRegisterService(repository Repository) *service {
	return &service{repository: repository}
}
func (s *service) RegisterService(input *InputRegister) (*models.EntityUser, string) {

	user := models.EntityUser{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&user)

	return resultRegister, errRegister
}
