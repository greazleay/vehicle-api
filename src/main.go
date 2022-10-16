package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/routes"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {

	router := gin.Default()

	routes.IndexRoutes(router)

	routes.MakeRoutes(router)

	routes.VehicleRoutes(router)

	router.Run()
}
