package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Visitor, Welcome to Vehicle API",
		})
	})

	router.Run()
}
