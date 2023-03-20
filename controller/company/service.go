package company

import "crud_api_company/model"

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAllCompany() ([]model.Company, string) {

	result, err := s.repository.GetAllCompany()

	return result, err
}
