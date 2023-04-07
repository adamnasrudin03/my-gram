package main

import (
	"fmt"
	"net/http"

	"github.com/adamnasrudin03/my-gram/app"
	"github.com/adamnasrudin03/my-gram/app/configs"
	"github.com/adamnasrudin03/my-gram/app/controller"
	routers "github.com/adamnasrudin03/my-gram/app/router"
	"github.com/adamnasrudin03/my-gram/pkg/database"
	"github.com/adamnasrudin03/my-gram/pkg/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "github.com/adamnasrudin03/my-gram/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	db *gorm.DB = database.SetupDbConnection()

	repo     = app.WiringRepository(db)
	services = app.WiringService(repo)

	userController controller.UserController = controller.NewUserController(services)
)

// @title My Gram API
// @version 1.0
// @description Service to manage MyGram data
// @termsOfService https://google.com
// @contact.name API Support
// @contact.email admin@mail.me
// @lisence.name Apache 2.0
// @lisence.url https://google.com
// @host localhost:8000
// @BasePath /
func main() {
	defer database.CloseDbConnection(db)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.APIResponse("welcome its server", http.StatusOK, "success", nil))
	})

	// Route here
	routers.UserRouter(router, userController)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helpers.APIResponse("page not found", http.StatusNotFound, "error", nil))
	})

	config := configs.GetInstance()
	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
