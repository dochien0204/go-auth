package employee

import (
	"crud_api_company/config"
	"crud_api_company/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	Database *gorm.DB
}

type Repository interface {
	GetAllEmployee() ([]model.Employee, string)
	AddEmployeeIntoCompany(companyId int, input *model.Employee) (*model.Employee, string)
	DeleteEmployee(employeeId int) (*model.Employee, string)
	GetAllEmployeeFromCompany(companyId int) ([]model.Employee, string)
}

func NewEmployeeRepository() *repository {
	return &repository{
		Database: config.DB,
	}
}

func (r *repository) GetAllEmployee() ([]model.Employee, string) {
	var employees []model.Employee

	result := r.Database.Find(&employees)

	if result.Error != nil {
		return nil, "Something went wrong"
	}

	return employees, "Successfully"
}

func (r *repository) AddEmployeeIntoCompany(companyId int, input *model.Employee) (*model.Employee, string) {

	result := r.Database.Create(&input)

	if result.Error != nil {
		return nil, "Something went wrong"
	}

	return input, "Successfully"

}

func (r *repository) DeleteEmployee(employeeId int) (*model.Employee, string) {
	var employee model.Employee
	result := r.Database.Table("employees").Clauses(clause.Returning{}).Unscoped().Delete(&employee, employeeId)

	if result.Error != nil {
		return nil, result.Error.Error()
	}

	return &employee, "Successfully"
}

func (r *repository) GetAllEmployeeFromCompany(companyId int) ([]model.Employee, string) {
	var employees []model.Employee

	result := r.Database.Where("company_id = ? ", companyId).Find(&employees)

	if result.Error != nil {
		return nil, result.Error.Error()
	}

	return employees, "Successfully"
}
