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

func GetAllVehicles(context *gin.Context) {

	var vehicles []models.Vehicle
	config.DB.Find(&vehicles)

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "All Vehicles",
		"data":       vehicles,
	})
}

func GetVehicleByID(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

	var vehicle models.Vehicle
	if err := config.DB.Preload("Make").First(&vehicle, "id = ?", params.ID).Error; err != nil {

		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"statusText": "failure",
			"statusCode": 404,
			"errorType":  "NotFoundException",
			"error":      err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "Vehicle with ID: " + params.ID,
		"data":       vehicle,
	})
}

func GetVehicleByModel(context *gin.Context) {

	// Validate Request Params
	params := dtos.VehicleModelDto{}

	if err := context.BindUri(&params); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

	var vehicle models.Vehicle
	if err := config.DB.Preload("Make").First(&vehicle, "model = ?", params.Model).Error; err != nil {

		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"statusText": "failure",
			"statusCode": 404,
			"errorType":  "NotFoundException",
			"error":      err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "Vehicle with Model: " + params.Model,
		"data":       vehicle,
	})
}

func UpdateVehicle(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}
	if err := context.BindUri(&params); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

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
	if err := config.DB.Where("model = ?", body.Model).Not("id = ?", params.ID).First(&vehicleExists).Error; err == nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"statusText": "failed",
			"statusCode": 409,
			"errorType":  "ConflictException",
			"error":      "Vehicle with model: " + body.Model + " already exists",
		})
		return
	}

	// Check if Vehicle to update exists
	var vehicle models.Vehicle
	result := config.DB.First(&vehicle, "id = ?", params.ID)

	if result.Error != nil {

		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"statusText": "failure",
			"statusCode": 404,
			"errorType":  "NotFoundException",
			"error":      result.Error.Error(),
		})
		return
	}

	config.DB.Model(&vehicle).Updates(models.Vehicle{
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
	})

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Make Updated",
		"data":       vehicle,
	})

}

func DeleteVehicle(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

	result := config.DB.Delete(&models.Vehicle{}, "id = ?", params.ID)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Vehicle Deleted",
	})
}
