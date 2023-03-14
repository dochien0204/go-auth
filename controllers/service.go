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

func (s *service) GetAllUser() ([]models.Users, string) {

	resultGetAll, errGetAll := s.repository.GetAllUser()

	return resultGetAll, errGetAll
}

func (s *service) GetUserById(userId uint) (*models.Users, string) {
	resultGet, errGet := s.repository.GetUserById(userId)

	return resultGet, errGet
}

func (s *service) UpdateUserById(userId uint, input *UserUpdateInput) (*models.Users, string) {
	var user models.Users

	user.Password = input.Password
	user.Email = input.Email

	resultUpdate, errUpdate := s.repository.UpdateUserById(userId, &user)

	return resultUpdate, errUpdate
}

func (s *service) DeleteUserById(userId uint) (*models.Users, string) {
	result, err := s.repository.DeleteUserById(userId)

	return result, err
}
