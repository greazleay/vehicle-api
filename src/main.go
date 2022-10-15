package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/controllers"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.GET("/", controllers.Index)

	makeRouter := router.Group("/v1/makes")

	makeRouter.POST("/", controllers.CreateMake)
	makeRouter.GET("/", controllers.GetAllMakes)
	makeRouter.GET("/:id", controllers.GetMakeByID)
	makeRouter.PATCH("/:id", controllers.UpdateMake)
	makeRouter.DELETE("/:id", controllers.DeleteMake)

	vehicleRouter := router.Group("/v1/vehicles")

	vehicleRouter.POST("/", controllers.CreateVehicle)

	router.Run()
}
