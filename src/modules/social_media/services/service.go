package services

import "insta-clone/src/modules/social_media/repositories"

type service struct {
	repository repositories.RepositorySocialMediaCommand
}

func NewSocialMediaService() *service {
	repository := repositories.NewSocialMediaRepository()
	service := &service{repository}

	return service
}
