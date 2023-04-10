package models

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint
	Message string
}
