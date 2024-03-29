package entities

import "insta-clone/src/modules/comment/entities"

// Photo represents the model for an photo
type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~title can't empty!"`
	Caption  string `gorm:"not null" json:"caption" form:"caption" valid:"required~caption can't empty!"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo url can't empty!"`
	UserID   uint
	ImageID  uint
	Comments []entities.Comment
}
