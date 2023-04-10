package models

type Photo struct {
	GormModel
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint
}
