package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/exceptions"
	"github.com/greazleay/vehicle-api/src/helpers"
	"github.com/greazleay/vehicle-api/src/models"
)

// CreateVehicle godoc
// @Summary      creates a new vehicle
// @Description  create vehicle
// @Tags         Vehicle
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param 		 data	body	dtos.CreateVehicleDto	true	"New Vehicle Details JSON"
// @Success      201  {object}  dtos.SuccessResponseDto{data=models.Vehicle}
// @Failure      400  {object}  dtos.FailedResponseDto	"request body validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      409  {object}  dtos.FailedResponseDto	"another vehicle with the same model property exists"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /vehicles [post]
func CreateVehicle(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateVehicleDto{}

	if err := context.BindJSON(&body); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Check if Vehicle with specified model already exists
	var vehicleExists models.Vehicle
	if err := config.DB.First(&vehicleExists, "model = ?", body.Model).Error; err == nil {

		exceptions.HandleConflictException(context, "Vehicle with model: "+body.Model+" already exists")
		return
	}

	var make models.Make
	if err := config.DB.First(&make, "id = ?", body.MakeID).Error; err != nil {
		exceptions.HandleNotFoundException(context, err)
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
		exceptions.HandleBadRequestException(context, result.Error)
		return
	}

	helpers.HandleCreatedResponse(context, "Vehicle Created", newVehicle)
}

// GetAllVehicles godoc
// @Summary      returns all vehicles
// @Description  get vehicles
// @Tags         Vehicle
// @Security 	JWT
// @Accept       json
// @Produce      json
// @success 200 {object} dtos.SuccessResponseDto{data=[]models.Vehicle} "all vehicles returned"
// @Failure      400  {object}  dtos.FailedResponseDto	"token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /vehicles [get]
func GetAllVehicles(context *gin.Context) {

	var allVehicles []models.Vehicle
	config.DB.Find(&allVehicles)

	helpers.HandleOkResponse(context, "All Vehicles", allVehicles)
}

// GetVehicleByID godoc
// @Summary      returns a vehicle by its 16 caharcter uuid
// @Description  get vehicle by ID
// @Tags         Vehicle
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Vehicle ID(UUID)"
// @success 200 {object} dtos.SuccessResponseDto{data=models.Vehicle} "vehicle with the specified ID returned"
// @Failure      400  {object}  dtos.FailedResponseDto	"request param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"vehicle with the specified ID not found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /vehicles/{id} [get]
func GetVehicleByID(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	var vehicle models.Vehicle
	if err := config.DB.Preload("Make").First(&vehicle, "id = ?", params.ID).Error; err != nil {

		exceptions.HandleNotFoundException(context, err)
		return
	}

	helpers.HandleOkResponse(context, "Vehicle with ID: "+params.ID, vehicle)
}

// GetVehicleByModel godoc
// @Summary      returns a vehicle by model
// @Description  get vehicle by model
// @Tags         Vehicle
// @Security 	 JWT
// @Accept       json
// @Produce      json
// @Param        model   query      string  true  "vehcile search by model"
// @success 	 200 {object} dtos.SuccessResponseDto{data=models.Vehicle} "vehicle with the search model returned"
// @Failure      400  {object}  dtos.FailedResponseDto	"request query validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"no vehicle with the specified model found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /vehicles/models [get]
func GetVehicleByModel(context *gin.Context) {

	// Validate Request Query
	query := dtos.VehicleModelDto{}

	if err := context.BindQuery(&query); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	var vehicle models.Vehicle
	if err := config.DB.Preload("Make").First(&vehicle, "model = ?", query.Model).Error; err != nil {

		exceptions.HandleNotFoundException(context, err)
		return
	}

	helpers.HandleOkResponse(context, "Vehicle with Model: "+query.Model, vehicle)
}

// UpdateVehicle godoc
// @Summary      updates a vehicle
// @Description  update vehicle
// @Tags         Vehicle
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Vehicle ID(UUID)"
// @Param 		 data	body	dtos.CreateVehicleDto	true	"Vehicle Details JSON"
// @success 200 {object} dtos.SuccessResponseDto{data=models.Vehicle} "vehicle updated suucessfully"
// @Failure      400  {object}  dtos.FailedResponseDto	"request body/param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"vehicle with ID in request params not found"
// @Failure      409  {object}  dtos.FailedResponseDto	"another vehicle with model in request body exists"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /vehicles/{id} [patch]
func UpdateVehicle(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}
	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Validate Request Body
	body := dtos.CreateVehicleDto{}
	if err := context.BindJSON(&body); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Check if Vehicle with specified model already exists
	var vehicleExists models.Vehicle
	if err := config.DB.Where("model = ?", body.Model).Not("id = ?", params.ID).First(&vehicleExists).Error; err == nil {

		exceptions.HandleConflictException(context, "Vehicle with model: "+body.Model+" already exists")
		return
	}

	// Check if Vehicle to update exists
	var vehicle models.Vehicle
	result := config.DB.First(&vehicle, "id = ?", params.ID)

	if result.Error != nil {

		exceptions.HandleNotFoundException(context, result.Error)
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

	helpers.HandleOkResponse(context, "Vehicle Updated", vehicle)

}

// DeleteVehicle godoc
// @Summary      deletes a vehicle
// @Description  delete vehicle
// @Tags         Vehicle
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Vehicle ID(UUID)"
// @success 200 {object} dtos.SuccessResponseDto "vehicle deleted successfully"
// @Failure      400  {object}  dtos.FailedResponseDto	"request param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /vehicles/{id} [delete]
func DeleteVehicle(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	result := config.DB.Delete(&models.Vehicle{}, "id = ?", params.ID)

	if result.Error != nil {
		exceptions.HandleBadRequestException(context, result.Error)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Vehicle Deleted",
	})
}
