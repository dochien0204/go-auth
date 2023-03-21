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

func (s *service) AddEmployeeIntoCompany(companyId int, input *EmployeeCreateInput) (*model.Employee, string) {

	var employee model.Employee
	employee.CompanyId = companyId
	employee.EmployeeAge = input.EmployeeAge
	employee.EmployeeName = input.EmployeeName
	employee.EmployeePhone = input.EmployeeName
	result, err := s.repository.AddEmployeeIntoCompany(companyId, &employee)

	return result, err
}

func (s *service) DeleteEmployee(employeeId int) (*model.Employee, string) {
	result, err := s.repository.DeleteEmployee(employeeId)

	return result, err
}

func (s *service) GetAllEmployeeFromCompany(companyId int) ([]model.Employee, string) {
	result, err := s.repository.GetAllEmployeeFromCompany(companyId)

	return result, err
}
