package handlers

import (
	"insta-clone/internals/utils"
	"insta-clone/src/modules/comment/dto"
	"net/http"
	"strconv"

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

func (h httpHandlerImpl) DeleteComment(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	err := h.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted:": true})
}

func (h httpHandlerImpl) GetOneComment(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	comment, err := h.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data one": comment})
}

func (h httpHandlerImpl) UpdateComment(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	commentRequestBody := dto.CommentRequestBody{}

	contentType := utils.GetContentType(ctx)

	if contentType == appJson {
		ctx.ShouldBindJSON(&commentRequestBody)
	} else {
		ctx.ShouldBind(&commentRequestBody)
	}

	comment, err := h.Update(id, commentRequestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}
