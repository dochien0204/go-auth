package department

import "crud_api_company/model"

type IDepartment interface {
	GetAllDepartMent() ([]model.Department, string)
	GetDepartmentById(departmentId int) (*model.Department, string)
	CreateNewDepartmetForCompany(input *DepartmentCreateInput) (*model.Department, string)
	UpdateDepartmentById(departmentId int, input *DepartmentUpdateInput) (*model.Department, string)
	DeleteDepartmentById(departmentId int) (*model.Department, string)
	GetAllEmployeeInDepartment(department int) ([]*model.Employee, string)
	DeleteEmployeeInDepartment(employeeId int, departmentId int) (*model.EmployeeDepartment, string)
	AddEmployeeIntoDepartment(employeeId int, departmentId int) (*model.EmployeeDepartment, string)
}
