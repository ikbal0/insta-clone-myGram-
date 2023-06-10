package repositories

import "insta-clone/src/modules/photo/entities"

type RepositoryPhotoCommand interface {
	Input(photo entities.Photo) (entities.Photo, error)
	Update(id int, photo entities.Photo) (entities.Photo, error)
	Delete(photo entities.Photo) error
	GetByID(id int) (entities.Photo, error)
	GetAll() ([]entities.Photo, error)
}

func (r repository) Input(photo entities.Photo) (entities.Photo, error) {
	err := r.db.Debug().Create(&photo).Error

	return photo, err
}

func (r *repository) Update(id int, photo entities.Photo) (entities.Photo, error) {
	var input entities.Photo

	input.Title = photo.Title
	input.Caption = photo.Caption

	err := r.db.Debug().Model(&photo).Where("Id = ?", id).Updates(&input).Error

	return photo, err
}

func (r *repository) Delete(photo entities.Photo) error {
	err := r.db.Debug().Delete(photo).Error

	return err
}

func (r *repository) GetAll() ([]entities.Photo, error) {
	var photo []entities.Photo
	err := r.db.Find(&photo).Error

	return photo, err
}

func (r *repository) GetByID(id int) (entities.Photo, error) {
	var photo entities.Photo
	err := r.db.First(&photo, "Id = ?", id).Error

	return photo, err
}
