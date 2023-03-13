package getallusers

import "jwt-project/models"

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAllUsers() ([]models.EntityUser, bool) {

	result, err := s.repository.GetAllUsers()

	if !err {
		return nil, false
	}

	return result, true
}
