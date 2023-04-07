package router

import (
	"github.com/adamnasrudin03/my-gram/app/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine, userController controller.UserController) {
	userRoutes := e.Group("/api/v1/auth")
	{
		// Create
		userRoutes.POST("/register", userController.Register)
		// Create
		userRoutes.POST("/login", userController.Login)
	}
}
