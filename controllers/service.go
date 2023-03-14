package controllers

import "crud-api/models"

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateNewUser(input *UserCreateInput) (*models.Users, string) {
	var user models.Users

	user.Username = input.Username
	user.Password = input.Password
	user.Email = input.Email
	resultCreate, errCreate := s.repository.CreateNewUser(&user)

	return resultCreate, errCreate

}

func (s *service) GetAllUser() ([]models.Users, string)            { return nil, "string" }
func (s *service) GetUserById(userId uint) (*models.Users, string) { return nil, "string" }
func (s *service) UpdateUserById(userId uint, input *UserUpdateInput) (*models.Users, string) {
	return nil, "string"
}
func (s *service) DeleteUserById(userId uint) (*models.Users, string) { return nil, "string" }
