package handlers

import (
	commentService "insta-clone/src/modules/comment/services"
	photoService "insta-clone/src/modules/photo/services"
	socialMediaService "insta-clone/src/modules/social_media/services"
	userService "insta-clone/src/modules/user/services"
)

type httpHandlerImpl struct {
	commentService.CommentService
	photoService.PhotoService
	socialMediaService.SocialMediaService
	userService.UserService
}

func NewHttpHandler() *httpHandlerImpl {
	commentService := commentService.NewCommentService()
	photoService := photoService.NewPhotoService()
	socialMediaService := socialMediaService.NewSocialMediaService()
	userService := userService.NewUserService()

	return &httpHandlerImpl{
		CommentService:     commentService,
		PhotoService:       photoService,
		SocialMediaService: socialMediaService,
		UserService:        userService,
	}
}
