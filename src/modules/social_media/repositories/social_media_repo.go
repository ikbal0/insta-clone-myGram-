package repositories

import (
	"insta-clone/src/modules/social_media/entities"
)

type RepositorySocialMediaCommand interface {
	Input(socialMedia entities.SocialMedia) (entities.SocialMedia, error)
	Update(id int, requestBody entities.SocialMedia) (entities.SocialMedia, error)
	Delete(socialMedia entities.SocialMedia) error
	GetByID(id int) (entities.SocialMedia, error)
	GetAll() ([]entities.SocialMedia, error)
}

func (r *repository) Input(socialMedia entities.SocialMedia) (entities.SocialMedia, error) {
	err := r.db.Debug().Create(&socialMedia).Error

	return socialMedia, err
}

func (r *repository) Update(id int, requestBody entities.SocialMedia) (entities.SocialMedia, error) {
	var socialMedia entities.SocialMedia
	r.db.First(&socialMedia, "Id = ?", id)
	err := r.db.Debug().Model(&socialMedia).Where("Id = ?", id).Updates(&requestBody).Error

	return socialMedia, err
}

func (r *repository) Delete(socialMedia entities.SocialMedia) error {
	err := r.db.Debug().Delete(socialMedia).Error

	return err
}

func (r *repository) GetAll() ([]entities.SocialMedia, error) {
	var socialMedia []entities.SocialMedia
	err := r.db.Find(&socialMedia).Error

	return socialMedia, err
}

func (r *repository) GetByID(id int) (entities.SocialMedia, error) {
	var socialMedia entities.SocialMedia
	err := r.db.First(&socialMedia, "Id = ?", id).Error

	return socialMedia, err
}
