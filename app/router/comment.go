package router

import (
	"adamnasrudin03/my-gram/app/controller"
	"adamnasrudin03/my-gram/app/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRouter(e *gin.Engine, h controller.CommentController) {
	commentRoutes := e.Group("/api/v1/comments")
	{
		commentRoutes.Use(middlewares.Authentication())
		commentRoutes.POST("/", h.CreateComment)
		commentRoutes.GET("/", middlewares.ListAuthorization(), h.GetAll)
		commentRoutes.PUT("/:id", h.UpdateComment)
		commentRoutes.GET("/:id", h.GetOne)
		commentRoutes.DELETE("/:id", h.DeleteComment)
	}
}
