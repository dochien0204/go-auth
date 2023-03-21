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
	result, err := s.repository.GetAllDepartMent()

	return result, err
}

func (s *service) CreateNewDepartmetForCompany(input *DepartmentCreateInput) (*model.Department, string) {

	var newDepartment model.Department
	newDepartment.DepartmentName = input.DepartmentName
	newDepartment.CompanyId = uint(input.CompanyId)
	result, err := s.repository.CreateNewDepartmetForCompany(&newDepartment)

	return result, err
}

func (s *service) GetDepartmentById(departmentId int) (*model.Department, string) {
	result, err := s.repository.GetDepartmentById(departmentId)

	return result, err
}

func (s *service) UpdateDepartmentById(departmentId int, input *DepartmentUpdateInput) (*model.Department, string) {
	var department model.Department

	department.CompanyId = uint(input.CompanyId)
	department.DepartmentName = input.DepartmentName

	result, err := s.repository.UpdateDepartmentById(departmentId, &department)

	return result, err
}

func (s *service) DeleteDepartmentById(departmentId int) (*model.Department, string) {
	result, err := s.repository.DeleteDepartmentById(departmentId)

	return result, err
}

func (s *service) GetAllEmployeeInDepartment(department int) ([]*model.Employee, string) {
	result, err := s.repository.GetAllEmployeeInDepartment(department)

	return result, err
}
