package routers

import (
	"insta-clone/controllers"
	"insta-clone/internals/protocols/http/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", controllers.Register)
		userRoute.POST("/login", controllers.Login)
	}

	socialMediaRoute := router.Group("/social-media")
	{
		socialMediaRoute.Use(middleware.Authentication())
		socialMediaRoute.POST("/", controllers.PostSocialMed)
		socialMediaRoute.GET("/", controllers.GetAllSocialMed)
		socialMediaRoute.GET("/:id", controllers.GetOneSocialMed)
		socialMediaRoute.PATCH("/:id", controllers.UpdateSocialMed)
		socialMediaRoute.DELETE("/:id", middleware.SocialMedAuthorization(), controllers.DeleteSocialMed)
	}

	commentRoute := router.Group("/comment")
	{
		commentRoute.Use(middleware.Authentication())
		commentRoute.POST("/", controllers.PostComment)
		commentRoute.PATCH("/:id", controllers.UpdateComment)
		commentRoute.GET("/", controllers.GetAllComment)
		commentRoute.GET("/:id", controllers.GetOneComment)
		commentRoute.DELETE("/:id", controllers.DeleteComment)
	}

	// router.GET("/test", controllers.GetAllPhoto)
	// router.POST("/test", controllers.PostTest)

	photoRoute := router.Group("/photo")
	{
		photoRoute.Use(middleware.Authentication())
		photoRoute.GET("/", controllers.GetAllPhoto)
		photoRoute.GET("/:photoId", controllers.GetOnePhoto)
	}

	router.MaxMultipartMemory = 8 << 20
	router.POST("/photo", middleware.Authentication(), controllers.UploadFile)
	router.DELETE("/photo/:photoId", controllers.DeleteImage)
	router.PATCH("/photo/:photoId", controllers.UpdatePhoto)

	return router
}
