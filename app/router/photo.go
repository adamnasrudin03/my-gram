package router

import (
	"adamnasrudin03/my-gram/app/controller"
	"adamnasrudin03/my-gram/app/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(e *gin.Engine, h controller.PhotoController) {
	photoRoutes := e.Group("/api/v1/photos")
	{
		photoRoutes.Use(middlewares.Authentication())
		photoRoutes.POST("/", h.CreatePhoto)
		photoRoutes.GET("/", middlewares.ListAuthorization(), h.GetAll)
		photoRoutes.PUT("/:id", h.UpdatePhoto)
		photoRoutes.GET("/:id", h.GetOne)
		photoRoutes.DELETE("/:id", h.DeletePhoto)
	}
}
