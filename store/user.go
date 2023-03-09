package store

type User struct {
	Username           string `json:"username"`
	Password           string `password:"password"`
	PasswordResetToken string `json:"password"`
}

var Users = []User{
	{Username: "dochien0204", Password: "chien123", PasswordResetToken: "abcxyz"},
	{Username: "dochien123", Password: "chien123", PasswordResetToken: "cbaxzy"},
}
