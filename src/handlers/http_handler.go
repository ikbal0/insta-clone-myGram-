package handlers

import "insta-clone/src/modules/comment/services"

type httpHandlerImpl struct {
	services.CommentService
}

func NewHttpHandler(commentService services.CommentService) *httpHandlerImpl {
	return &httpHandlerImpl{
		CommentService: commentService,
	}
}
