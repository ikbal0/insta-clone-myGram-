package models

type SocialMedia struct {
	GormModel
	Name           string
	SocialMediaUrl string
	UserID         uint
}
