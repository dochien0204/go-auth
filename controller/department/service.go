package department

import "crud_api_company/model"

type service struct {
	repository IRepository
}

func NewDepartmentService(repository IRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAllDepartMent() ([]model.Department, string) {
	result, err := s.GetAllDepartMent()

	return result, err
}
