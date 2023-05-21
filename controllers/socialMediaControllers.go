package controllers

import (
	"insta-clone/database"
	"insta-clone/internals/utils"
	"insta-clone/src/modules/social_media/entities"
	"net/http"

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

func UpdateSocialMed(c *gin.Context) {
	var db = database.GetDB()

	var socialMed entities.SocialMedia

	err := db.First(&socialMed, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	var input entities.SocialMedia

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&socialMed).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": socialMed})
}

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

func PostSocialMed(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	db := database.GetDB()
	contentType := utils.GetContentType(ctx)
	_, _ = db, contentType
	SocialMedia := entities.SocialMedia{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

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
