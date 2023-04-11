package main

import (
	"fmt"
	"net/http"

	"adamnasrudin03/my-gram/app"
	"adamnasrudin03/my-gram/app/configs"
	"adamnasrudin03/my-gram/app/controller"
	routers "adamnasrudin03/my-gram/app/router"
	"adamnasrudin03/my-gram/pkg/database"
	"adamnasrudin03/my-gram/pkg/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"adamnasrudin03/my-gram/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	db *gorm.DB = database.SetupDbConnection()

	repo     = app.WiringRepository(db)
	services = app.WiringService(repo)

	userController        controller.UserController        = controller.NewUserController(services)
	socialMediaController controller.SocialMediaController = controller.NewSocialMediaController(services)
	commentController     controller.CommentController     = controller.NewCommentController(services)
	photoController       controller.PhotoController       = controller.NewPhotoController(services)
)

// @title           MyGram Documentation
// @version         1.0
// @description     This is REST API documentation it's the final project of the training program (Scalable Web Services with Golang) organized by Kominfo x Hacktiv8

// @contact.name   Adam Nasrudin
// @contact.url    https://www.linkedin.com/in/adam-nasrudin/
// @contact.email  adamnasrudin858@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Add "Bearer" followed by a space and JWT token.
func main() {
	defer database.CloseDbConnection(db)
	config := configs.GetInstance()
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", config.Appconfig.Port)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.APIResponse("welcome its server", http.StatusOK, "false"))
	})

	// Route here
	routers.UserRouter(router, userController)
	routers.SocialMediaRouter(router, socialMediaController)
	routers.CommentRouter(router, commentController)
	routers.PhotoRouter(router, photoController)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helpers.APIResponse("page not found", http.StatusNotFound, "error"))
	})

	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
