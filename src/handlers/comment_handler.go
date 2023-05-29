package handlers

import (
	"insta-clone/database"
	"insta-clone/internals/utils"
	"insta-clone/src/modules/comment/dto"
	"insta-clone/src/modules/comment/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h httpHandlerImpl) PostComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	requestBody := dto.CommentRequestBody{}
	contentType := utils.GetContentType(ctx)

	if contentType == appJson {
		ctx.ShouldBindJSON(&requestBody)
	} else {
		ctx.ShouldBind(&requestBody)
	}

	requestBody.UserID = userID
	requestBody.PhotoID = 13

	comment, err := h.Input(requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}

func (h httpHandlerImpl) GetAllComment(ctx *gin.Context) {
	comments, err := h.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comments})
}

func DeleteComment(ctx *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var comment entities.Comment
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

	var comment []entities.Comment

	err := db.First(&comment, "Id = ?", ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data one": comment})
}

func UpdateComment(ctx *gin.Context) {
	var db = database.GetDB()

	var comment entities.Comment

	err := db.First(&comment, "Id = ?", ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	var input entities.Comment

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&comment).Updates(&input)

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}

func PostComment(ctx *gin.Context) {
	db := database.GetDB()
	contentType := utils.GetContentType(ctx)
	_, _ = db, contentType
	Comment := entities.Comment{}

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
