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

	router.GET("/test", controllers.GetAllPhoto)
	router.POST("/test", controllers.PostTest)

	router.MaxMultipartMemory = 8 << 20
	router.POST("/photo", controllers.UploadFile)

	return router
}
