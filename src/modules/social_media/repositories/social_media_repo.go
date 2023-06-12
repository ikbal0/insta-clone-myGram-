package repositories

import "insta-clone/src/modules/social_media/entities"

type RepositorySocialMediaCommand interface {
	Input(socialMedia entities.SocialMedia) (entities.SocialMedia, error)
	Update(id int, socialMedia entities.SocialMedia) (entities.SocialMedia, error)
	Delete(socialMedia entities.SocialMedia) error
	GetByID(id int) (entities.SocialMedia, error)
	GetAll() ([]entities.SocialMedia, error)
}
