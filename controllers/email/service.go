package email

import (
	"github.com/go-mail/mail"
)

type email struct {
}

func NewMail() *email {
	return &email{}
}

func (e *email) SendSimpleMessage(email string, message string) {
	m := mail.NewMessage()

	m.SetHeader("From", "doxuanchienh02042002@gmail.com")

	m.SetHeader("To", "chien.doxuan@vti.com.vn")

	m.SetAddressHeader("Cc", "doxuanchienh02042002@gmaiil.com", "Oliver")

	m.SetHeader("Subject", "RESET PASSWORD!")

	m.SetBody("text/plain", message)

	d := mail.NewDialer("smtp.gmail.com", 587, "doxuanchienh02042002@gmail.com", "gggrltyekbsxrhma")

	// Send the email to Kate, Noah and Oliver.

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}
}
