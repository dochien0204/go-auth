package login

import "jwt-project/models"

type Service interface {
	LoginService(input *InputLogin) (*models.EntityUser, string)
}
