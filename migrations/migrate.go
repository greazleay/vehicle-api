package main

import (
	"github.com/greazleay/vehicle-api/initializers"
	"github.com/greazleay/vehicle-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Maker{}, &models.Vehicle{})
}
