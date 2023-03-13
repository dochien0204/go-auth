package login

import (
	"jwt-project/configs"
	"jwt-project/models"

	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(input *models.EntityUser) (*models.EntityUser, string)
}

type repository struct {
	database *gorm.DB
}

func NewLoginRepository() *repository {
	return &repository{
		database: configs.Database,
	}
}

func (r *repository) LoginRepository(input *models.EntityUser) (*models.EntityUser, string) {
	var user models.EntityUser

	result := r.database.Table("entity_users").Where("username = ? AND password = ?", input.Username, input.Password).Find(&user)

	if err := result.Error; err != nil {
		return nil, "Error in db"
	}
	if result.RowsAffected == 0 {
		return nil, "Account not exists"
	}

	return &user, "Successfully"

}
