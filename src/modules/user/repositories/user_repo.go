package repositories

import (
	"insta-clone/src/modules/user/entities"
)

type RepositoryUserCommand interface {
	Input(user entities.User) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
}

func (r *repository) Input(user entities.User) (entities.User, error) {
	err := r.db.Debug().Create(&user).Error

	return user, err
}

func (r *repository) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	err := r.db.First(&user, "email = ?", email).Error

	return user, err
}
