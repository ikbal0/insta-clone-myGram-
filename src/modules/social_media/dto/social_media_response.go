package dto

import "time"

type SocialMediaUpdateResponse struct {
	Name           string     `json:"name" form:"name" valid:"required~name is required and can't be empty"`
	SocialMediaUrl string     `json:"social_media_url" form:"social_media_url" valid:"required~social media url is required and can't be empty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}
