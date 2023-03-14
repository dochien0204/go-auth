package controllers

import (
	"crud-api/configs"
	"crud-api/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateNewUser(input *models.Users) (*models.Users, string)
	GetAllUser() ([]models.Users, string)
	GetUserById(userId uint) (*models.Users, string)
	UpdateUserById(userId uint, input *models.Users) (*models.Users, string)
	DeleteUserById(userId uint) (*models.Users, string)
}

func NewRepository() *repository {
	return &repository{
		db: configs.DB,
	}
}

func (r *repository) CreateNewUser(input *models.Users) (*models.Users, string) {

	//check exists user
	rowCount := r.db.Where("username = ? ", input.Username).Find(&input).RowsAffected

	if rowCount != 0 {
		return nil, "User already exists"
	}

	result := r.db.Create(&input)

	if result.Error != nil {
		return nil, "Something went wrong"
	}

	return input, "Successfully"
}

func (r *repository) GetAllUser() ([]models.Users, string) {
	return nil, ""
}
func (r *repository) GetUserById(userId uint) (*models.Users, string) {
	return nil, ""

}
func (r *repository) UpdateUserById(userId uint, input *models.Users) (*models.Users, string) {
	return nil, ""

}
func (r *repository) DeleteUserById(userId uint) (*models.Users, string) {
	return nil, ""

}
