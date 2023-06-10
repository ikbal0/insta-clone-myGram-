package handlers

import (
	commentService "insta-clone/src/modules/comment/services"
	photoService "insta-clone/src/modules/photo/services"
)

type httpHandlerImpl struct {
	commentService.CommentService
	photoService.PhotoService
}

func NewHttpHandler() *httpHandlerImpl {
	commentService := commentService.NewCommentService()
	photoService := photoService.NewPhotoService()

	return &httpHandlerImpl{
		CommentService: commentService,
		PhotoService:   photoService,
	}
}
