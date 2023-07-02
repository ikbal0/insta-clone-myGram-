package handlers

import (
	"insta-clone/internals/utils"
	"insta-clone/src/modules/comment/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Post Comment godoc
// @Summary Post comment for a given id
// @Description Create Comment corresponding to the photo id in param
// @Tags Comment
// @Accept json
// @Produce json
// @Param dto.CommentRequestBody body dto.CommentRequestBody true "create comment"
// @Success 200 {object} entities.Comment
// @Router /comment [post]
// @Security Bearer
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

	comment, err := h.CommentService.Input(requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}

// Get All Comment godoc
// @Summary Get details
// @Description Get details of all comment
// @Tags Comment
// @Accept json
// @Produce json
// @Success 200 {object} entities.Comment
// @Router /comment [get]
// @Security Bearer
func (h httpHandlerImpl) GetAllComment(ctx *gin.Context) {
	comments, err := h.CommentService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comments})
}

// Delete Comment godoc
// @Summary Delete comment identified by given id
// @Description Delete the comment corresponding to the input id
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment to be deleted"
// @Success 204 "No content"
// @Router /comment/{Id} [delete]
// @Security Bearer
func (h httpHandlerImpl) DeleteComment(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	err := h.CommentService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted:": true})
}

// Get Comment By Id godoc
// @Summary Get details for a given id
// @Description Get details of comment corresponding to the input id
// @Tags Comment
// @Accept json
// @Produce json
// @Success 200 {object} entities.Comment
// @Router /comment/{Id} [get]
// @Security Bearer
func (h httpHandlerImpl) GetOneComment(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	comment, err := h.CommentService.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data one": comment})
}

// Update Comment godoc
// @Summary Update comment identified by given id
// @Description Update details of Comment corresponding to the input id
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment to be updated"
// @Success 200 {object} entities.Comment
// @Router /comment/{Id} [patch]
// @Security Bearer
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

	comment, err := h.CommentService.Update(id, commentRequestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}
