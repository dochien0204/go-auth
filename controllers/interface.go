package controllers

import "crud-api/models"

type IUserService interface {
	CreateNewUser(input *UserCreateInput) (*models.Users, string)
	GetAllUser() ([]models.Users, string)
	GetUserById(userId uint) (*models.Users, string)
	UpdateUserById(userId uint, input *UserUpdateInput) (*models.Users, string)
	DeleteUserById(userId uint) (*models.Users, string)
}
