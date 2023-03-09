package register

import "jwt-project/models"

type IService interface {
	RegisterService(input *InputRegister) (*models.EntityUser, string)
}
