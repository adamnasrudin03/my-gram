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
		commentRoutes.POST("/", middlewares.CheckAuthorization(), h.CreateComment)
		commentRoutes.GET("/", middlewares.CheckAuthorization(), h.GetAll)
		commentRoutes.PUT("/:id", middlewares.CommentAuthorization(), h.UpdateComment)
		commentRoutes.GET("/:id", middlewares.CheckAuthorization(), h.GetOne)
		commentRoutes.DELETE("/:id", middlewares.CommentAuthorization(), h.DeleteComment)
	}
}
