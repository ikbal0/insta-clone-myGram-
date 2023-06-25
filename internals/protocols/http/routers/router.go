package routers

import (
	"insta-clone/internals/protocols/http/middleware"
	"insta-clone/src/handlers"

	_ "insta-clone/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Insta Clone Api
// @version 1.0
// @description This is a simple services for managing cars
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email devyad@gmail.com
// @license.name Apace 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	router := gin.Default()
	handler := handlers.NewHttpHandler()

	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", handler.Register)
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
		// Create
		commentRoute.POST("/", handler.PostComment)
		// Update
		commentRoute.PATCH("/:id", handler.UpdateComment)
		// Read All
		commentRoute.GET("/", handler.GetAllComment)
		// Read
		commentRoute.GET("/:id", handler.GetOneComment)
		// Delete
		commentRoute.DELETE("/:id", handler.DeleteComment)
	}

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
