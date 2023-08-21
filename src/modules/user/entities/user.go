package entities

import (
	"errors"
	"insta-clone/internals/utils"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	GormModel
	Username string `gorm:"not null; uniqueIndex" json:"username" form:"username" valid:"required~username is required and can't be empty"`
	Email    string `gorm:"not null; uniqueIndex" json:"email" form:"email" valid:"required~email is required and can't be empty, email~Format email Invalid!"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required, minstringlength(8)~Password has to have a minimum length of 8 characters"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Age is required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age < 9 {
		return errors.New("come back when u at least 9 years old")
	}

	u.Password = utils.HashPass(u.Password)

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
