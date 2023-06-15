package repositories

import (
	"insta-clone/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewSocialMediaRepository() *repository {
	db := database.GetDB()
	tx := &repository{db: db}
	return tx
}
