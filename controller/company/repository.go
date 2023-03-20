package company

import (
	"crud_api_company/config"
	"crud_api_company/model"
	"crud_api_company/utils"

	"gorm.io/gorm"
)

type repository struct {
	Database *gorm.DB
}

type Repository interface {
	GetAllCompany() ([]model.Company, string)
}

func NewRepository() *repository {
	return &repository{
		Database: config.DB,
	}
}

func (r *repository) GetAllCompany() ([]model.Company, string) {
	var companies []model.Company

	result := r.Database.Preload("Departments").Find(&companies)

	resultCheck, errString := utils.CheckError(result)

	if !resultCheck {
		return nil, errString
	}

	return companies, errString

}
