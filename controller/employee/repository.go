package employee

import (
	"crud_api_company/config"
	"crud_api_company/model"

	"gorm.io/gorm"
)

type repository struct {
	Database *gorm.DB
}

type Repository interface {
	GetAllEmployee() ([]model.Employee, string)
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
