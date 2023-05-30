package services

import "insta-clone/src/modules/comment/repositories"

type service struct {
	repository repositories.RepositoryCommentCommand
}

func NewCommentService() *service {
	repository := repositories.NewCommentRepository()
	service := &service{repository}

	return service
}
