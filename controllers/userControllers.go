package controllers

import (
	"insta-clone/database"
	"insta-clone/helpers"
	"insta-clone/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func Register(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"username":   User.Username,
		"email":      User.Email,
		"password":   User.Password,
		"age":        User.Age,
		"created_at": User.CreatedAt,
		"updated_at": User.UpdatedAt,
	})
}
