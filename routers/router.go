package routers

import (
	"insta-clone/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", controllers.Register)
	}

	router.GET("/view", controllers.GetAllPhoto)

	return router
}
