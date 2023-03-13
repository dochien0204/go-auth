package getallusers

import "jwt-project/models"

type Service interface {
	GetAllUsers() ([]models.EntityUser, bool)
}
