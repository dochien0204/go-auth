package employee

import "crud_api_company/model"

type service struct {
	repository Repository
}

func NewEmployeeService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAllEmployee() ([]model.Employee, string) {
	result, err := s.repository.GetAllEmployee()

	return result, err
}
