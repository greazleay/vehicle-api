package main

import (
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Make{}, &models.Vehicle{}, &models.User{})
}
