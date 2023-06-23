package handlers

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

func (h httpHandlerImpl) Login(c *gin.Context) {
	contentType := utils.GetContentType(c)
	User := entities.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	data, err := h.UserService.GetByEmail(User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": data,
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": data,
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
