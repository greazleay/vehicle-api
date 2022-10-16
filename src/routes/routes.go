package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/controllers"
)

func IndexRoutes(indexRouter *gin.Engine) {
	indexRouter.GET("/", controllers.Index)
}

func MakeRoutes(router *gin.Engine) {

	makeRouter := router.Group("/v1/makes")

	{
		makeRouter.POST("/", controllers.CreateMake)
		makeRouter.GET("/", controllers.GetAllMakes)
		makeRouter.GET("/:id", controllers.GetMakeByID)
		makeRouter.PATCH("/:id", controllers.UpdateMake)
		makeRouter.DELETE("/:id", controllers.DeleteMake)
	}

}

func VehicleRoutes(router *gin.Engine) {

	vehicleRouter := router.Group("/v1/vehicles")

	{
		vehicleRouter.POST("/", controllers.CreateVehicle)
	}

}
