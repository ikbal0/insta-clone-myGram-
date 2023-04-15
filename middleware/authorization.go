package middleware

import (
	"fmt"
	"insta-clone/database"
	"insta-clone/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SocialMedAuthorization() gin.HandlerFunc {
	fmt.Println("social media Author")
	return func(ctx *gin.Context) {
		db := database.GetDB()
		SocialMedId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("UserID").First(&SocialMedia, uint(SocialMedId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data didn't exist",
			})
			return
		}

		if SocialMedia.UserID == userID {
			ctx.Next()
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		}
	}
}
