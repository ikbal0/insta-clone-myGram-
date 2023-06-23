package routers

import (
	"insta-clone/internals/protocols/http/middleware"
	"insta-clone/src/handlers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()
	handler := handlers.NewHttpHandler()

	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", handlers.Register)
		userRoute.POST("/login", handler.Login)
	}

	socialMediaRoute := router.Group("/social-media")
	{
		socialMediaRoute.Use(middleware.Authentication())
		socialMediaRoute.POST("/", handler.PostSocialMed)
		socialMediaRoute.GET("/", handler.GetAllSocialMed)
		socialMediaRoute.GET("/:id", handler.GetOneSocialMed)
		socialMediaRoute.PATCH("/:id", handler.UpdateSocialMed)
		socialMediaRoute.DELETE("/:id", middleware.SocialMedAuthorization(), handler.DeleteSocialMed)
	}

	commentRoute := router.Group("/comment")
	{
		commentRoute.Use(middleware.Authentication())
		commentRoute.POST("/", handler.PostComment)
		commentRoute.PATCH("/:id", handler.UpdateComment)
		commentRoute.GET("/", handler.GetAllComment)
		commentRoute.GET("/:id", handler.GetOneComment)
		commentRoute.DELETE("/:id", handler.DeleteComment)
	}

	// router.GET("/test", controllers.GetAllPhoto)
	// router.POST("/test", controllers.PostTest)

	photoRoute := router.Group("/photo")
	{
		photoRoute.Use(middleware.Authentication())
		photoRoute.GET("/", handler.GetAllPhoto)
		photoRoute.GET("/:photoId", handler.GetOnePhoto)
	}

	router.MaxMultipartMemory = 8 << 20
	router.POST("/photo", middleware.Authentication(), handler.UploadFile)
	router.DELETE("/photo/:photoId", handler.DeleteImage)
	router.PATCH("/photo/:photoId", handler.UpdatePhoto)

	return router
}
