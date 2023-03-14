package getallusers

import (
	"errors"
	"jwt-project/configs"
	"jwt-project/models"

	"gorm.io/gorm"
)

type repository struct {
	database *gorm.DB
}

type Repository interface {
	GetAllUsers() ([]models.EntityUser, bool)
	GetUserByUserName(username string) (*models.EntityUser, string)
	UpdateUserById(userId uint, input *models.EntityUser) (*models.EntityUser, string)
	DeleteUserById(userId uint) (*models.EntityUser, string)
}

func NewRepository() *repository {
	return &repository{
		database: configs.Database,
	}
}

func (r *repository) GetAllUsers() ([]models.EntityUser, bool) {
	var accounts []models.EntityUser
	err := r.database.Table("entity_users").Find(&accounts).Error

	if err != nil {
		return nil, false
	}
	return accounts, true
}

func (r *repository) GetUserByUserName(username string) (*models.EntityUser, string) {
	var user models.EntityUser

	result := r.database.Where("username = ? ", username).Find(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, "Not found user has username = " + username
	}

	return &user, "Successfully"
}

func (r *repository) UpdateUserById(userId uint, input *models.EntityUser) (*models.EntityUser, string) {
	var user models.EntityUser
	result := r.database.Find(&user, userId)

	if result.RowsAffected == 0 {
		return nil, "User not exists"
	}

	user.Username = input.Username
	user.Password = input.Password
	user.PasswordResetToken = input.PasswordResetToken
	user.Email = input.Email

	r.database.Save(&user)

	return &user, "Successfully"

}

func (r *repository) DeleteUserById(userId uint) (*models.EntityUser, string) {
	var user models.EntityUser
	result := r.database.Find(&user, userId)

	if result.RowsAffected == 0 {
		return nil, "User does not exists"
	}

	sqlStatement := `DELETE FROM entity_users where id = $1;`

	r.database.Exec(sqlStatement, userId)

	return &user, "Successfully"
}
