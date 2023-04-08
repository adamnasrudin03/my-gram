package router

import (
	"adamnasrudin03/my-gram/app/controller"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(e *gin.Engine, h controller.SocialMediaController) {
	socialMediaRoutes := e.Group("/api/v1/social-media")
	{
		socialMediaRoutes.POST("/", h.CreateSocialMedia)
		socialMediaRoutes.GET("/", h.GetAll)
		socialMediaRoutes.PUT("/:id", h.UpdateSocialMedia)
		socialMediaRoutes.GET("/:id", h.GetOne)
		socialMediaRoutes.DELETE("/:id", h.DeleteSocialMedia)
	}
}
