package forgot

import (
	"jwt-project/models"
	"jwt-project/store"
)

type repository struct{}

type Repository interface {
	ForgotRepository(input *models.EntityUser) (*models.EntityUser, string)
}

func NewForgotRepository() *repository {
	return &repository{}
}

func (r *repository) ForgotRepository(input *models.EntityUser) (*models.EntityUser, string) {

	var user models.EntityUser

	isExists := false
	for _, val := range store.Users {
		if val.Username == input.Username {
			user.Username = val.Username
			user.Password = val.Password
			user.PasswordResetToken = val.PasswordResetToken
			isExists = true
			break
		}
	}

	if !isExists {
		return &user, "Account not exists"
	}

	if user.PasswordResetToken != input.PasswordResetToken {
		return nil, "Token invalid"
	}

	return &user, "Forgot successfully"

}
