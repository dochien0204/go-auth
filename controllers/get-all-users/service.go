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

func (s *service) GetUserByUserName(username string) (*models.EntityUser, string) {
	result, err := s.repository.GetUserByUserName(username)
	return result, err
}
