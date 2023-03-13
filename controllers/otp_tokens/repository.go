package otptokens

import (
	"jwt-project/configs"
	"jwt-project/models"
	"time"

	"gorm.io/gorm"
)

type IOTPRepository interface {
	GenerateNewOTP(token string) *models.OTPToken
}

type repository struct {
	database *gorm.DB
}

func NewRepository() *repository {
	return &repository{
		database: configs.Database,
	}
}

func (r *repository) GenerateNewOTP(token string) *models.OTPToken {
	var otpToken models.OTPToken

	otpToken.ExpiredAt = time.Now().Add(time.Minute * 10)
	otpToken.Token = token
	otpToken.ID = 0

	r.database.Create(&otpToken)
	return &otpToken

}
