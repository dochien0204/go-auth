package login

import "jwt-project/models"

type Service interface {
	LoginService(input *InputLogin) (*models.EntityUser, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *InputLogin) (*models.EntityUser, string) {
	user := models.EntityUser{
		Username: input.Username,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin

}
