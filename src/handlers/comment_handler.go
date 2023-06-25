package handlers

import (
	"insta-clone/internals/utils"
	"insta-clone/src/modules/comment/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// CreateCars godoc
// @Summary Post details for a given id
// @Description Get details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param entities.Comment body entities.Comment true "create car"
// @Success 200 {object} entities.Comment
// @Router /cars [post]
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

// GetAllCar godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} entities.Comment
// @Router /cars [get]
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

// DeleteCar godoc
// @Summary Delete car identified by given id
// @Description Delete the car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No content"
// @Router /cars/{Id} [delete]
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

// GetOneCars godoc
// @Summary Get details for a given id
// @Description Get details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} entities.Comment
// @Router /cars/{Id} [get]
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

// UpdateCar godoc
// @Summary Update car identified by given id
// @Description Update details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} entities.Comment
// @Router /cars/{Id} [patch]
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
