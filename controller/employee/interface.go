package employee

import "crud_api_company/model"

type IEmployee interface {
	GetAllEmployee() ([]model.Employee, string)
	AddEmployeeIntoCompany(companyId int, input *EmployeeCreateInput) (*model.Employee, string)
	DeleteEmployee(employeeId int) (*model.Employee, string)
	GetAllEmployeeFromCompany(companyId int) ([]model.Employee, string)
}
