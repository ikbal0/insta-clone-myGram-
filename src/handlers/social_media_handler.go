package handlers

import (
	"insta-clone/internals/utils"
	"insta-clone/src/modules/social_media/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h httpHandlerImpl) DeleteSocialMed(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	err := h.SocialMediaService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted:": true})
}

func (h httpHandlerImpl) UpdateSocialMed(ctx *gin.Context) {
	socialMedia := entities.SocialMedia{}
	getId := ctx.Param("id")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	contentType := utils.GetContentType(ctx)

	if contentType == appJson {
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}

	UpdatedSocialMedia, err := h.SocialMediaService.Update(id, socialMedia)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": UpdatedSocialMedia})
}

func (h httpHandlerImpl) GetOneSocialMed(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	socialMedia, err := h.SocialMediaService.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "record has not found!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data one": socialMedia})
}

func (h httpHandlerImpl) GetAllSocialMed(ctx *gin.Context) {
	socialMedia, err := h.SocialMediaService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": socialMedia})
}

func (h httpHandlerImpl) PostSocialMed(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := utils.GetContentType(ctx)
	SocialMedia := entities.SocialMedia{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia, err := h.SocialMediaService.Input(userID, SocialMedia)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post Social Med",
		"data":    SocialMedia,
	})
}
