package handlers

import (
	"insta-clone/database"
	"insta-clone/internals/utils"
	"insta-clone/src/modules/social_media/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func DeleteSocialMed(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var socialMedDelete entities.SocialMedia
	err := db.First(&socialMedDelete, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	db.Delete(&socialMedDelete)

	c.JSON(http.StatusOK, gin.H{"deleted:": true})
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

// func UpdateSocialMed(c *gin.Context) {
// 	var db = database.GetDB()

// 	var socialMed entities.SocialMedia

// 	err := db.First(&socialMed, "Id = ?", c.Param("id")).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
// 		return
// 	}

// 	var input entities.SocialMedia

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db.Model(&socialMed).Updates(&input)

// 	c.JSON(http.StatusOK, gin.H{"data": socialMed})
// }

func GetOneSocialMed(c *gin.Context) {
	var db = database.GetDB()

	var socialMedOne []entities.SocialMedia

	err := db.First(&socialMedOne, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data one": socialMedOne})
}

func GetAllSocialMed(ctx *gin.Context) {
	var db = database.GetDB()

	var socialMed []entities.SocialMedia

	err := db.Find(&socialMed).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": socialMed})
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
