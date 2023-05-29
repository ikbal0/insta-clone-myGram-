package services

import "insta-clone/src/modules/comment/repositories"

type service struct {
	repository repositories.RepositoryCommentCommand
}

func NewService(repository repositories.RepositoryCommentCommand) *service {
	service := &service{repository}

	return service
}
