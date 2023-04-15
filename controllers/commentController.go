package controllers

import (
	"insta-clone/database"
	"insta-clone/helpers"
	"insta-clone/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteComment(ctx *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var comment models.Comment
	err := db.First(&comment, "Id = ?", ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	db.Delete(&comment)

	ctx.JSON(http.StatusOK, gin.H{"deleted:": true})
}

func GetOneComment(ctx *gin.Context) {
	var db = database.GetDB()

	var comment []models.Comment

	err := db.First(&comment, "Id = ?", ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data one": comment})
}

func GetAllComment(ctx *gin.Context) {
	var db = database.GetDB()

	var comment []models.Comment

	err := db.Find(&comment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}

func UpdateComment(ctx *gin.Context) {
	var db = database.GetDB()

	var comment models.Comment

	err := db.First(&comment, "Id = ?", ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	var input models.Comment

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&comment).Updates(&input)

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}

func PostComment(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	Comment := models.Comment{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	Comment.UserID = 1
	Comment.PhotoID = 13

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Comment posted",
		"data":    Comment,
	})
}
