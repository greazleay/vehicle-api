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

	makerRouter := router.Group("/v1/makers")

	makerRouter.POST("/", controllers.CreateMaker)
	makerRouter.GET("/", controllers.GetAllMakers)
	makerRouter.GET("/:id", controllers.GetMakerByID)
	makerRouter.PATCH("/:id", controllers.UpdateMaker)
	makerRouter.DELETE("/:id", controllers.DeleteMaker)

	vehicleRouter := router.Group("/v1/vehicles")

	vehicleRouter.POST("/", controllers.CreateVehicle)

	router.Run()
}
