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
}

func NewDepartmentRepository() *repository {
	return &repository{
		Database: config.DB,
	}
}

func (r *repository) GetAllDepartMent() ([]model.Department, string) {
	var departments []model.Department

	result := r.Database.Raw("Select * from departments").Scan(&departments)

	resultCheck, errCheck := utils.CheckError(result)

	if !resultCheck {
		return nil, errCheck
	}

	return departments, errCheck
}
