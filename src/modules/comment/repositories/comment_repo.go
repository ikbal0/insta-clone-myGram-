package repositories

import (
	"insta-clone/src/modules/comment/entities"
)

type RepositoryCommentCommand interface {
	Input(comment entities.Comment) (entities.Comment, error)
	Update(id int, comment entities.Comment) (entities.Comment, error)
	Delete(comment entities.Comment) error
	GetByID(id int) (entities.Comment, error)
	GetAll() ([]entities.Comment, error)
}

func (r *repository) Input(comment entities.Comment) (entities.Comment, error) {
	err := r.db.Debug().Create(&comment).Error

	return comment, err
}

func (r *repository) Update(id int, comment entities.Comment) (entities.Comment, error) {
	var input entities.Comment

	input.PhotoID = comment.PhotoID
	input.UserID = comment.UserID
	input.Message = comment.Message

	err := r.db.Debug().Model(&comment).Where("Id = ?", id).Updates(&input).Error

	return comment, err
}

func (r *repository) Delete(comment entities.Comment) error {
	err := r.db.Debug().Delete(comment).Error

	return err
}

func (r *repository) GetAll() ([]entities.Comment, error) {
	var comment []entities.Comment
	err := r.db.Find(&comment).Error

	return comment, err
}

func (r *repository) GetByID(id int) (entities.Comment, error) {
	var comment entities.Comment
	err := r.db.First(&comment, "Id = ?", id).Error

	return comment, err
}
