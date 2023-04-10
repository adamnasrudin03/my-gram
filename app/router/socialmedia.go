package router

import (
	"adamnasrudin03/my-gram/app/controller"
	"adamnasrudin03/my-gram/app/middlewares"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(e *gin.Engine, h controller.SocialMediaController) {
	socialMediaRoutes := e.Group("/api/v1/social-media")
	{
		socialMediaRoutes.Use(middlewares.Authentication())
		socialMediaRoutes.POST("/", middlewares.CheckAuthorization(), h.CreateSocialMedia)
		socialMediaRoutes.GET("/", middlewares.CheckAuthorization(), h.GetAll)
		socialMediaRoutes.PUT("/:id", middlewares.SocialMediaAuthorization(), h.UpdateSocialMedia)
		socialMediaRoutes.GET("/:id", middlewares.CheckAuthorization(), h.GetOne)
		socialMediaRoutes.DELETE("/:id", middlewares.SocialMediaAuthorization(), h.DeleteSocialMedia)
	}
}
