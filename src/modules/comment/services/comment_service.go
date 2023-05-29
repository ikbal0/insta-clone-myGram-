package services

import (
	"insta-clone/src/modules/comment/dto"
	"insta-clone/src/modules/comment/entities"
)

type CommentService interface {
	Input(data dto.CommentRequestBody) (entities.Comment, error)
	Update(id int, data dto.CommentRequestBody) (entities.Comment, error)
	GetByID(id int) (entities.Comment, error)
	GetAll() ([]entities.Comment, error)
	Delete(id int) error
}
