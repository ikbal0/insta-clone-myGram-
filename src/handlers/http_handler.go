package handlers

import (
	commentService "insta-clone/src/modules/comment/services"
	photoService "insta-clone/src/modules/photo/services"
	socialMediaService "insta-clone/src/modules/social_media/services"
)

type httpHandlerImpl struct {
	commentService.CommentService
	photoService.PhotoService
	socialMediaService.SocialMediaService
}

func NewHttpHandler() *httpHandlerImpl {
	commentService := commentService.NewCommentService()
	photoService := photoService.NewPhotoService()
	socialMediaService := socialMediaService.NewSocialMediaService()

	return &httpHandlerImpl{
		CommentService:     commentService,
		PhotoService:       photoService,
		SocialMediaService: socialMediaService,
	}
}
