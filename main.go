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
)

var (
	db *gorm.DB = database.SetupDbConnection()

	repo     = app.WiringRepository(db)
	services = app.WiringService(repo)

	userController controller.UserController = controller.NewUserController(services)
)

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

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helpers.APIResponse("page not found", http.StatusNotFound, "error", nil))
	})

	config := configs.GetInstance()
	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
