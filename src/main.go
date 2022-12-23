package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	_ "github.com/greazleay/vehicle-api/src/docs"
	"github.com/greazleay/vehicle-api/src/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

// @title           Vehicle API
// @version         1.0
// @description     This is a simple CRUD API for managing vehicle information.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Lekan Adetunmbi
// @contact.url    https://qxz.netlify.app
// @contact.email  lothbroch@gmail.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      api-vehicle.onrender.com
// @BasePath  /v1
// @schemes   http https

// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {

	router := gin.Default()

	routes.IndexRoutes(router)

	routes.MakeRoutes(router)

	routes.VehicleRoutes(router)

	routes.UserRoutes(router)

	routes.AuthRoutes((router))

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
