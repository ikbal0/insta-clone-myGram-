package controllers

import (
	"insta-clone/database"
	"insta-clone/internals/utils"
	"insta-clone/src/modules/user/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	User := entities.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := utils.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "Invalid email/password",
		})
		return
	}

	token := utils.TokenGenerator(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Register(ctx *gin.Context) {
	db := database.GetDB()
	contentType := utils.GetContentType(ctx)
	_, _ = db, contentType
	User := entities.User{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"username":   User.Username,
		"email":      User.Email,
		"password":   User.Password,
		"age":        User.Age,
		"created_at": User.CreatedAt,
		"updated_at": User.UpdatedAt,
	})
}
