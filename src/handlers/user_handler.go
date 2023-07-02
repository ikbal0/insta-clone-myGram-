package handlers

import (
	"insta-clone/internals/utils"
	"insta-clone/src/modules/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param dto.UserLoginRequestBody body dto.UserLoginRequestBody true "login"
// @Success 200 {object} "token"
// @Router /user/login [post]
func (h httpHandlerImpl) Login(c *gin.Context) {
	contentType := utils.GetContentType(c)
	User := dto.UserLoginRequestBody{}

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
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": data,
	})
}

// Register User godoc
// @Summary Register User
// @Description Register User
// @Tags User
// @Accept json
// @Produce json
// @Param dto.UserRequestBody body dto.UserRequestBody true "user register"
// @Success 200 {object} entities.User
// @Router /user/register [post]
func (h httpHandlerImpl) Register(ctx *gin.Context) {
	contentType := utils.GetContentType(ctx)

	data := dto.UserRequestBody{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&data)
	} else {
		ctx.ShouldBind(&data)
	}

	// err := db.Debug().Create(&User).Error
	user, err := h.UserService.Input(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
