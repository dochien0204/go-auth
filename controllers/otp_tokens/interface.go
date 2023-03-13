package otptokens

import "jwt-project/models"

type IOtp interface {
	GenerateNewOTP(token string) *models.OTPToken
}
