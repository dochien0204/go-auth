package controllers

type UserCreateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserUpdateInput struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
