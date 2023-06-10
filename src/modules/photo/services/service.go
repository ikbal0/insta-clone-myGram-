package services

import "insta-clone/src/modules/photo/repositories"

type service struct {
	repository repositories.RepositoryPhotoCommand
}

func NewPhotoService() *service {
	repository := repositories.NewPhotoRepository()
	service := &service{repository}

	return service
}
