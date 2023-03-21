package department

import (
	"crud_api_company/config"
	"crud_api_company/model"
	"crud_api_company/utils"

	"gorm.io/gorm"
)

type repository struct {
	Database *gorm.DB
}

type IRepository interface {
	GetAllDepartMent() ([]model.Department, string)
	GetDepartmentById(departmentId int) (*model.Department, string)
	CreateNewDepartmetForCompany(input *model.Department) (*model.Department, string)
	UpdateDepartmentById(departmentId int, input *model.Department) (*model.Department, string)
	DeleteDepartmentById(departmentId int) (*model.Department, string)
	GetAllEmployeeInDepartment(department int) ([]*model.Employee, string)
}

func NewDepartmentRepository() *repository {
	return &repository{
		Database: config.DB,
	}
}

func (r *repository) GetAllDepartMent() ([]model.Department, string) {
	var departments []model.Department

	result := r.Database.Find(&departments)

	resultCheck, errCheck := utils.CheckError(result)

	if !resultCheck {
		return nil, errCheck
	}

	return departments, errCheck
}

func (r *repository) CreateNewDepartmetForCompany(input *model.Department) (*model.Department, string) {
	//check exsist department in company
	rowCount := r.Database.Where("department_name = ? and company_id = ?", input.DepartmentName, input.CompanyId).Find(&input).RowsAffected

	if rowCount != 0 {
		return nil, input.DepartmentName + " is exists in company"
	}

	result := r.Database.Create(&input)

	if result.Error != nil {
		return nil, "Something went wrong"
	}

	return input, "Successfully"

}

func (r *repository) GetDepartmentById(departmentId int) (*model.Department, string) {

	var department model.Department

	result := r.Database.Find(&department, departmentId)

	if result.Error != nil {
		return nil, result.Error.Error()
	}

	if result.RowsAffected == 0 {
		return nil, "Department doesn't exists"
	}

	return &department, "Successfully"
}

func (r *repository) UpdateDepartmentById(departmentId int, input *model.Department) (*model.Department, string) {
	var department model.Department

	result := r.Database.Find(&department, departmentId)

	if result.Error != nil {
		return nil, result.Error.Error()
	}

	if result.RowsAffected == 0 {
		return nil, "Department doesn't exists"
	}

	if r.Database.Find(&model.Company{}, input.CompanyId).RowsAffected == 0 {
		return nil, "Company not exists"
	}

	department.CompanyId = input.CompanyId
	department.DepartmentName = input.DepartmentName

	r.Database.Save(&department)

	return &department, "Successfully"
}

func (r *repository) DeleteDepartmentById(departmentId int) (*model.Department, string) {
	var department model.Department
	result := r.Database.Unscoped().Delete(&department, departmentId)

	if result.Error != nil {
		return nil, "Cannot delete for some problems"
	}

	if result.RowsAffected == 0 {
		return nil, "Department doesn't exists"
	}

	return &department, "Successfully"
}

func (r *repository) GetAllEmployeeInDepartment(department int) ([]*model.Employee, string) {
	var departments model.Department

	err := r.Database.Preload("Employees").Find(&departments, department)

	if err.Error != nil {
		return nil, err.Error.Error()
	}

	return departments.Employees, "Successfully"
}
