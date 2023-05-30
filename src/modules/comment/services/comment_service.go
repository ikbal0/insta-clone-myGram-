package services

import (
	"insta-clone/src/modules/comment/dto"
	"insta-clone/src/modules/comment/entities"
)

type CommentService interface {
	Input(data dto.CommentRequestBody) (entities.Comment, error)
	Update(id int, data dto.CommentRequestBody) (dto.CommentResponse, error)
	GetByID(id int) (entities.Comment, error)
	GetAll() ([]entities.Comment, error)
	Delete(id int) error
}

func (s service) Input(data dto.CommentRequestBody) (entities.Comment, error) {
	comment := entities.Comment{
		UserID:  data.UserID,
		PhotoID: data.PhotoID,
		Message: data.Message,
	}

	newComment, err := s.repository.Input(comment)

	return newComment, err
}

func (s service) Update(id int, data dto.CommentRequestBody) (dto.CommentResponse, error) {
	comment := entities.Comment{
		UserID:  data.UserID,
		PhotoID: data.PhotoID,
		Message: data.Message,
	}

	newComment, err := s.repository.Update(id, comment)

	return newComment, err
}

func (s service) Delete(id int) error {
	comment, err := s.repository.GetByID(id)

	if err != nil {
		return err
	}

	errDel := s.repository.Delete(comment)

	return errDel
}

func (s service) GetAll() ([]entities.Comment, error) {
	comment, err := s.repository.GetAll()

	return comment, err
}

func (s service) GetByID(id int) (entities.Comment, error) {
	comment, err := s.repository.GetByID(id)

	return comment, err
}
