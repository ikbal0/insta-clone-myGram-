package routers

import (
	"insta-clone/internals/protocols/http/middleware"
	"insta-clone/src/handlers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", handlers.Register)
		userRoute.POST("/login", handlers.Login)
	}

	socialMediaRoute := router.Group("/social-media")
	{
		socialMediaRoute.Use(middleware.Authentication())
		socialMediaRoute.POST("/", handlers.PostSocialMed)
		socialMediaRoute.GET("/", handlers.GetAllSocialMed)
		socialMediaRoute.GET("/:id", handlers.GetOneSocialMed)
		socialMediaRoute.PATCH("/:id", handlers.UpdateSocialMed)
		socialMediaRoute.DELETE("/:id", middleware.SocialMedAuthorization(), handlers.DeleteSocialMed)
	}

	commentRoute := router.Group("/comment")
	{
		commentRoute.Use(middleware.Authentication())
		commentRoute.POST("/", handlers.PostComment)
		commentRoute.PATCH("/:id", handlers.UpdateComment)
		commentRoute.GET("/", handlers.GetAllComment)
		commentRoute.GET("/:id", handlers.GetOneComment)
		commentRoute.DELETE("/:id", handlers.DeleteComment)
	}

	// router.GET("/test", controllers.GetAllPhoto)
	// router.POST("/test", controllers.PostTest)

	photoRoute := router.Group("/photo")
	{
		photoRoute.Use(middleware.Authentication())
		photoRoute.GET("/", handlers.GetAllPhoto)
		photoRoute.GET("/:photoId", handlers.GetOnePhoto)
	}

	router.MaxMultipartMemory = 8 << 20
	router.POST("/photo", middleware.Authentication(), handlers.UploadFile)
	router.DELETE("/photo/:photoId", handlers.DeleteImage)
	router.PATCH("/photo/:photoId", handlers.UpdatePhoto)

	return router
}
