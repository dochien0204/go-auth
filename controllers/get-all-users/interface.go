package getallusers

import "jwt-project/models"

type Service interface {
	GetAllUsers() ([]models.EntityUser, bool)
	GetUserByUserName(username string) (*models.EntityUser, string)
	UpdateUserById(userId uint, input *InputUpdateUser) (*models.EntityUser, string)
	DeleteUserById(userId uint) (*models.EntityUser, string)
}
