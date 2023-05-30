package repositories

import "insta-clone/src/modules/photo/entities"

type RepositoryPhotoCommand interface {
	Input(comment entities.Photo) (entities.Photo, error)
	Update(id int, comment entities.Photo) (entities.Photo, error)
	Delete(comment entities.Photo) error
	GetByID(id int) (entities.Photo, error)
	GetAll() ([]entities.Photo, error)
}
