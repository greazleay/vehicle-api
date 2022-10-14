package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/controllers"
	"github.com/greazleay/vehicle-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.GET("/", controllers.Index)
	router.POST("/makers", controllers.CreateMaker)
	router.GET("/makers", controllers.GetAllMakers)
	router.GET("/makers/:id", controllers.GetMakerByID)
	router.PATCH("/makers/:id", controllers.UpdateMaker)
	router.DELETE("/makers/:id", controllers.DeleteMaker)

	router.Run()
}
