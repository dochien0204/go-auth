package getallusers

import "jwt-project/models"

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAllUsers() ([]models.EntityUser, bool) {

	result, err := s.repository.GetAllUsers()

	if !err {
		return nil, false
	}

	return result, true
}

func (s *service) GetUserByUserName(username string) (*models.EntityUser, string) {
	result, err := s.repository.GetUserByUserName(username)
	return result, err
}

func (s *service) UpdateUserById(userId uint, input *InputUpdateUser) (*models.EntityUser, string) {

	var user models.EntityUser
	user.Username = input.Username
	user.Password = input.Password
	user.PasswordResetToken = input.PasswordResetToken
	user.Email = input.Email
	resultUpdate, errUpdate := s.repository.UpdateUserById(userId, &user)

	return resultUpdate, errUpdate
}

func (s *service) DeleteUserById(userID uint) (*models.EntityUser, string) {

	resultDelete, errDelete := s.repository.DeleteUserById(userID)

	return resultDelete, errDelete
}
