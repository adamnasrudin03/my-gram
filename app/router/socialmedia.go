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
		socialMediaRoutes.POST("/", h.CreateSocialMedia)
		socialMediaRoutes.GET("/", middlewares.ListAuthorization(), h.GetAll)
		socialMediaRoutes.PUT("/:id", h.UpdateSocialMedia)
		socialMediaRoutes.GET("/:id", h.GetOne)
		socialMediaRoutes.DELETE("/:id", h.DeleteSocialMedia)
	}
}
