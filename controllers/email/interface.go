package email

type Service interface {
	SendSimpleMessage(email string, message string)
}
