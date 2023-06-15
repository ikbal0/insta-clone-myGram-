package services

import (
	"insta-clone/src/modules/social_media/dto"
	"insta-clone/src/modules/social_media/entities"
)

type SocialMediaService interface {
	Input(id uint, data entities.SocialMedia) (entities.SocialMedia, error)
	Update(id int, data entities.SocialMedia) (dto.SocialMediaUpdateResponse, error)
	GetByID(id int) (entities.SocialMedia, error)
	GetAll() ([]entities.SocialMedia, error)
	Delete(id int) error
}

func (s service) Input(id uint, data entities.SocialMedia) (entities.SocialMedia, error) {
	data.UserID = id
	socialMedia, err := s.repository.Input(data)

	return socialMedia, err
}

func (s service) Update(id int, data entities.SocialMedia) (dto.SocialMediaUpdateResponse, error) {
	socialMedia, err := s.repository.Update(id, data)

	socialMediaResponse := dto.SocialMediaUpdateResponse{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UpdatedAt:      socialMedia.UpdatedAt,
	}

	return socialMediaResponse, err
}

func (s service) Delete(id int) error {
	socialMedia, err := s.repository.GetByID(id)

	if err != nil {
		return err
	}

	errDel := s.repository.Delete(socialMedia)

	return errDel
}

func (s service) GetAll() ([]entities.SocialMedia, error) {
	socialMedia, err := s.repository.GetAll()

	return socialMedia, err
}

func (s service) GetByID(id int) (entities.SocialMedia, error) {
	socialMedia, err := s.repository.GetByID(id)

	return socialMedia, err
}
