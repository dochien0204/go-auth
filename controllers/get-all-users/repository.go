package getallusers

import (
	"jwt-project/configs"
	"jwt-project/models"

	"gorm.io/gorm"
)

type repository struct {
	Database *gorm.DB
}

type Repository interface {
	GetAllUsers() ([]models.EntityUser, bool)
}

func NewRepository() *repository {
	return &repository{
		Database: configs.Database,
	}
}

func (r *repository) GetAllUsers() ([]models.EntityUser, bool) {
	var accounts []models.EntityUser
	err := r.Database.Find(&accounts).Error

	if err != nil {
		return nil, false
	}
	return accounts, true
}
