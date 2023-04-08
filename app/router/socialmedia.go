package router

import (
	"adamnasrudin03/my-gram/app/controller"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(e *gin.Engine, h controller.SocialMediaController) {
	socialMediaRoutes := e.Group("/api/v1/social-media")
	{
		// Create
		socialMediaRoutes.POST("/", h.CreateSocialMedia)
		// Get
		socialMediaRoutes.GET("/", h.GetAll)
		// Get
		socialMediaRoutes.GET("/{id}", h.GetOne)
		// Update
		socialMediaRoutes.PUT("/{id}", h.UpdateSocialMedia)
	}
}
