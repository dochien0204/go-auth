package getallusers

import (
	"errors"
	"jwt-project/configs"
	"jwt-project/models"

	"gorm.io/gorm"
)

type repository struct {
	database *gorm.DB
}

type Repository interface {
	GetAllUsers() ([]models.EntityUser, bool)
	GetUserByUserName(username string) (*models.EntityUser, string)
}

func NewRepository() *repository {
	return &repository{
		database: configs.Database,
	}
}

func (r *repository) GetAllUsers() ([]models.EntityUser, bool) {
	var accounts []models.EntityUser
	err := r.database.Table("entity_users").Find(&accounts).Error

	if err != nil {
		return nil, false
	}
	return accounts, true
}

func (r *repository) GetUserByUserName(username string) (*models.EntityUser, string) {
	var user models.EntityUser

	result := r.database.Where("username = ? ", username).Find(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, "Not found user has username = " + username
	}

	return &user, "Successfully"
}
