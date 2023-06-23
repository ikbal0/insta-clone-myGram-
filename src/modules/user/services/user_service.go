package services

import (
	"errors"
	"insta-clone/internals/utils"
	"insta-clone/src/modules/user/dto"
	"insta-clone/src/modules/user/entities"
)

type UserService interface {
	Input(data dto.UserRequestBody) (entities.User, error)
	GetByEmail(formData entities.User) (string, error)
}

func (s service) Input(data dto.UserRequestBody) (entities.User, error) {
	user := entities.User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Age:      data.Age,
	}

	newUser, err := s.repository.Input(user)

	return newUser, err
}

func (s service) GetByEmail(formData entities.User) (string, error) {
	user, err := s.repository.GetByEmail(formData.Email)

	comparePass := utils.ComparePass([]byte(user.Password), []byte(formData.Password))

	if !comparePass {
		return "Invalid email/password", errors.New("Unauthorize")
	}

	token := utils.TokenGenerator(user.ID, user.Email)

	return token, err
}
