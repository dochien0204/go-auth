package register

import (
	"jwt-project/configs"
	"jwt-project/models"
	"jwt-project/utils"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *models.EntityUser) (*models.EntityUser, string)
}

type repository struct {
	database *gorm.DB
}

func NewRegisterRepository() *repository {
	return &repository{
		database: configs.Database,
	}
}

func (r *repository) RegisterRepository(input *models.EntityUser) (*models.EntityUser, string) {
	var user models.EntityUser

	//check user exists in db
	result := r.database.Where("username = ? ", input.Username).Find(&user)

	if result.Error != nil {
		return nil, "Error in DB"
	}

	if result.RowsAffected != 0 {
		return nil, "Account is exists"
	}

	user.Username = input.Username
	user.Password = input.Password
	user.PasswordResetToken = utils.GenerateSecureToken(3)
	r.database.Create(&user)
	return &user, "Successfully"

}
