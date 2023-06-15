package dto

type SocialMediaUpdateRequest struct {
	Name           string `json:"name" form:"name"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url"`
}
