package login

import (
	"jwt-project/models"
	"jwt-project/store"
)

type Repository interface {
	LoginRepository(input *models.EntityUser) (*models.EntityUser, string)
}

type repository struct {
}

func NewLoginRepository() *repository {
	return &repository{}
}

func (r *repository) LoginRepository(input *models.EntityUser) (*models.EntityUser, string) {
	var user models.EntityUser

	user.Username = input.Username
	user.Password = input.Password

	isExists := false
	for _, val := range store.Users {
		if val.Username == user.Username && val.Password == user.Password {
			isExists = true
			break
		}
	}

	if !isExists {
		return &user, "Invalid account"
	}

	return &user, "Login successfully"

}
