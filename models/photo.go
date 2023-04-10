package models

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~title can't empty!"`
	Caption  string `gorm:"not null" json:"caption" form:"caption" valid:"required~caption can't empty!"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo url can't empty!"`
	UserID   uint
}
