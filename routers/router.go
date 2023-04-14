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

	socialMediaRoute := router.Group("/social-media")
	{
		socialMediaRoute.POST("/", controllers.PostSocialMed)
		socialMediaRoute.GET("/", controllers.GetAllSocialMed)
		socialMediaRoute.GET("/:id", controllers.GetOneSocialMed)
		socialMediaRoute.PATCH("/:id", controllers.UpdateSocialMed)
		socialMediaRoute.DELETE("/:id", controllers.DeleteSocialMed)
	}

	// router.GET("/test", controllers.GetAllPhoto)
	// router.POST("/test", controllers.PostTest)

	router.MaxMultipartMemory = 8 << 20
	router.POST("/photo", controllers.UploadFile)
	router.DELETE("/photo/:photoId", controllers.DeleteImage)

	return router
}
