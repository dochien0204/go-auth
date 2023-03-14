package otptokens

import "jwt-project/models"

type IOtp interface {
	GenerateNewOTP(token string, user uint) *models.OTPToken
}
