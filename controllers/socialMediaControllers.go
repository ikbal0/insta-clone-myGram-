package controllers

import (
	"insta-clone/database"
	"insta-clone/helpers"
	"insta-clone/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSocialMed(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var socialMedDelete models.SocialMedia
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

	var socialMed models.SocialMedia

	err := db.First(&socialMed, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	var input models.SocialMedia

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&socialMed).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": socialMed})
}

func GetOneSocialMed(c *gin.Context) {
	var db = database.GetDB()

	var socialMedOne []models.SocialMedia

	err := db.First(&socialMedOne, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data one": socialMedOne})
}

func GetAllSocialMed(ctx *gin.Context) {
	var db = database.GetDB()

	var socialMed []models.SocialMedia

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
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = 1

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
