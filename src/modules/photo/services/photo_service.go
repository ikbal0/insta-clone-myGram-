package services

import (
	"insta-clone/src/modules/photo/entities"
)

type PhotoService interface {
	// Input(data dto.PhotoRequestBody) (entities.Photo, error)
	// Update(id int, data dto.PhotoRequestBody) (entities.Photo, error)
	Input(data entities.Photo) (entities.Photo, error)
	Update(id int, data entities.Photo) (entities.Photo, error)
	GetByID(id int) (entities.Photo, error)
	GetAll() ([]entities.Photo, error)
	Delete(id int) error
}

func (s service) Input(data entities.Photo) (entities.Photo, error) {
	// comment := entities.Comment{
	// 	UserID:  data.UserID,
	// 	PhotoID: data.PhotoID,
	// 	Message: data.Message,
	// }

	newPhoto, err := s.repository.Input(data)

	return newPhoto, err
}

func (s service) Update(id int, data entities.Photo) (entities.Photo, error) {
	// comment := entities.Comment{
	// 	UserID:  data.UserID,
	// 	PhotoID: data.PhotoID,
	// 	Message: data.Message,
	// }

	newPhoto, err := s.repository.Update(id, data)

	return newPhoto, err
}

func (s service) Delete(id int) error {
	photo, err := s.repository.GetByID(id)

	if err != nil {
		return err
	}

	errDel := s.repository.Delete(photo)

	return errDel
}

func (s service) GetAll() ([]entities.Photo, error) {
	photo, err := s.repository.GetAll()

	return photo, err
}

func (s service) GetByID(id int) (entities.Photo, error) {
	photo, err := s.repository.GetByID(id)

	return photo, err
}
