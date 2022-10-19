package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/controllers"
	"github.com/greazleay/vehicle-api/src/middlewares"
)

func IndexRoutes(indexRouter *gin.Engine) {
	indexRouter.GET("/", controllers.Index)
}

func MakeRoutes(router *gin.Engine) {

	makeRouter := router.Group("/v1/makes").Use(middlewares.Auth())

	{
		makeRouter.POST("/", controllers.CreateMake)
		makeRouter.GET("/", controllers.GetAllMakes)
		makeRouter.GET("/names", controllers.GetMakeByName)
		makeRouter.GET("/countries", controllers.GetMakesByCountry)
		makeRouter.GET("/:id", controllers.GetMakeByID)
		makeRouter.PATCH("/:id", controllers.UpdateMake)
		makeRouter.DELETE("/:id", controllers.DeleteMake)
	}

}

func VehicleRoutes(router *gin.Engine) {

	vehicleRouter := router.Group("/v1/vehicles").Use(middlewares.Auth())

	{
		vehicleRouter.POST("/", controllers.CreateVehicle)
		vehicleRouter.GET("/", controllers.GetAllVehicles)
		vehicleRouter.GET("/models", controllers.GetVehicleByModel)
		vehicleRouter.GET("/:id", controllers.GetVehicleByID)
		vehicleRouter.PATCH("/:id", controllers.UpdateVehicle)
		vehicleRouter.DELETE("/:id", controllers.DeleteVehicle)
	}

}

func UserRoutes(router *gin.Engine) {

	userRouter := router.Group("/v1/users")

	{
		userRouter.POST("/", controllers.CreateUser)
		userRouter.GET("/", middlewares.Auth(), controllers.GetAllUsers)
		userRouter.GET("/:id", middlewares.Auth(), controllers.GetUserByID)
		userRouter.PATCH("/:id", middlewares.Auth(), controllers.UpdateUser)
		userRouter.DELETE("/:id", middlewares.Auth(), controllers.DeleteUser)
	}
}

func AuthRoutes(router *gin.Engine) {

	authRouter := router.Group("/v1/auth")

	{
		authRouter.POST("/login", controllers.LoginUser)
	}
}
