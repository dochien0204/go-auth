package forgot

import (
	"fmt"
	"jwt-project/configs"
	"jwt-project/models"

	"gorm.io/gorm"
)

type repository struct {
	database *gorm.DB
}

type Repository interface {
	ForgotRepository(input *models.EntityUser) (*models.EntityUser, bool)
}

func NewForgotRepository() *repository {
	return &repository{
		database: configs.Database,
	}
}

func (r *repository) ForgotRepository(input *models.EntityUser) (*models.EntityUser, bool) {

	var user models.EntityUser

	countRow := r.database.Table("entity_users").Where("username = ? ", input.Username).Find(&user).RowsAffected

	if countRow == 0 {
		return nil, false
	}

	fmt.Println("Username", user.Username)

	return &user, true

}
