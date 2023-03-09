package register

import (
	"fmt"
	"jwt-project/models"
	"jwt-project/store"
	"jwt-project/utils"
)

type Repository interface {
	RegisterRepository(input *models.EntityUser) (*models.EntityUser, string)
}

type repository struct {
}

func NewRegisterRepository() *repository {
	return &repository{}
}

func (r *repository) RegisterRepository(input *models.EntityUser) (*models.EntityUser, string) {
	var user models.EntityUser
	user.Password = input.Password
	user.Username = input.Username
	user.PasswordResetToken = utils.GenerateSecureToken(6)
	checkExists := false
	userEntered := store.User{Username: input.Username, Password: input.Password, PasswordResetToken: input.PasswordResetToken}
	for _, val := range store.Users {
		if userEntered.Username == val.Username {
			checkExists = true
			break
		}
	}

	if checkExists {
		return nil, "Account already exists"
	}

	store.Users = append(store.Users, userEntered)
	fmt.Println(store.Users)

	return &user, "Register successfully"

}
