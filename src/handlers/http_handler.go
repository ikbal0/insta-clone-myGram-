package handlers

import "insta-clone/src/modules/comment/services"

type httpHandlerImpl struct {
	services.CommentService
}

func NewHttpHandler() *httpHandlerImpl {
	commentService := services.NewCommentService()

	return &httpHandlerImpl{
		CommentService: commentService,
	}
}
