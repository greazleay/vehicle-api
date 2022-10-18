package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/models"
)

func CreateVehicle(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateVehicleDto{}

	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

	// Check if Vehicle with specified model already exists
	var vehicleExists models.Vehicle
	if err := config.DB.First(&vehicleExists, "model = ?", body.Model).Error; err == nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"statusText": "failed",
			"statusCode": 409,
			"errorType":  "ConflictException",
			"error":      "Vehicle with model: " + body.Model + " already exists",
		})
		return
	}

	var make models.Make
	if err := config.DB.First(&make, "id = ?", body.MakeID).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"statusText": "failure",
			"statusCode": 404,
			"errorType":  "NotFoundException",
			"error":      err.Error() + ", Make with specified ID Not Found",
		})
		return
	}

	// Create new Vehicle
	newVehicle := models.Vehicle{
		Model:         body.Model,
		Category:      body.Category,
		Year:          body.Year,
		ImageUrl:      body.ImageUrl,
		Price:         body.Price,
		NumberOfSeats: body.NumberOfSeats,
		Cylinder:      body.Cylinder,
		HorsePower:    body.HorsePower,
		Torque:        body.Torque,
		TopSpeed:      body.TopSpeed,
		Acceleration:  body.Acceleration,
		Transmission:  body.Transmission,
	}

	newVehicle.Make = make

	result := config.DB.Create(&newVehicle)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"statusText": "success",
		"statusCode": 201,
		"message":    "Vehicle Created",
		"data":       newVehicle,
	})
}
