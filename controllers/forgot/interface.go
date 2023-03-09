package forgot

import (
	"jwt-project/models"
)

type IForgotService interface {
	ForgotService(input *InputForgot) (*models.EntityUser, string)
}
