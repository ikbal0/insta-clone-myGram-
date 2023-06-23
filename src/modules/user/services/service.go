package services

import "insta-clone/src/modules/user/repositories"

type service struct {
	repository repositories.RepositoryUserCommand
}

func NewUserService() *service {
	repository := repositories.NewUserRepository()
	service := &service{repository}

	return service
}
