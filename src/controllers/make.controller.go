package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/exceptions"
	"github.com/greazleay/vehicle-api/src/helpers"
	"github.com/greazleay/vehicle-api/src/models"
	"gorm.io/gorm"
)

// CreateMake godoc
// @Summary      creates a new make
// @Description  create make
// @Tags         Make
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param 		 data	body	dtos.CreateMakeDto	true	"Make Details JSON"
// @Success      201  {object}  dtos.SuccessResponseDto{data=models.Make}
// @Failure      400  {object}  dtos.FailedResponseDto	"request body validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      409  {object}  dtos.FailedResponseDto	"make with the same name exists"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /makes [post]
func CreateMake(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateMakeDto{}

	if err := context.BindJSON(&body); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Check if Make with specified name already exists
	var makeExists models.Make

	if err := config.DB.First(&makeExists, "name = ?", body.Name).Error; err == nil {

		exceptions.HandleConflictException(context, "Make with name: "+body.Name+" already exists")
		return
	}

	// Create new Make
	newMake := models.Make{Name: body.Name, Country: body.Country}

	result := config.DB.Create(&newMake)

	if result.Error != nil {
		exceptions.HandleBadRequestException(context, result.Error)
		return
	}

	helpers.HandleCreatedResponse(context, "Make Created", newMake)
}

// GetAllMakes godoc
// @Summary      returns all makes
// @Description  get makes
// @Tags         Make
// @Security 	JWT
// @Accept       json
// @Produce      json
// @success 200 {object} dtos.SuccessResponseDto{data=[]models.Make} "all makes returned"
// @Failure      400  {object}  dtos.FailedResponseDto	"token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /makes [get]
func GetAllMakes(context *gin.Context) {

	var allMakes []models.Make
	config.DB.Find(&allMakes)

	helpers.HandleOkResponse(context, "All Makes", allMakes)

}

// GetMakeByID godoc
// @Summary      returns a make by its 16 caharcter uuid
// @Description  get make by ID
// @Tags         Make
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Make ID(UUID)"
// @success 200 {object} dtos.SuccessResponseDto{data=models.Make} "make with the specified ID returned"
// @Failure      400  {object}  dtos.FailedResponseDto	"request param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"make with the specified ID not found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /makes/{id} [get]
func GetMakeByID(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	var make models.Make
	if err := config.DB.Preload("Vehicles", func(db *gorm.DB) *gorm.DB {
		return db.Select("MakeID", "Model")
	}).First(&make, "id = ?", params.ID).Error; err != nil {

		exceptions.HandleNotFoundException(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "Make with ID: " + params.ID,
		"data":       make,
	})
}

// GetMakeByName godoc
// @Summary      returns a make by name
// @Description  get make by name
// @Tags         Make
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        name   query      string  true  "make search by name"
// @success 200 {object} dtos.SuccessResponseDto{data=models.Make} "make with the search name returned"
// @Failure      400  {object}  dtos.FailedResponseDto	"request query validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"no make with the search name found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /makes/names [get]
func GetMakeByName(context *gin.Context) {

	// Validate Request Query
	query := dtos.MakeNameDto{}

	if err := context.BindQuery(&query); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	var make models.Make
	if err := config.DB.Preload("Vehicles", func(db *gorm.DB) *gorm.DB {
		return db.Select("MakeID", "Model")
	}).First(&make, "name = ?", query.Name).Error; err != nil {

		exceptions.HandleNotFoundException(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "Make with Name: " + query.Name,
		"data":       make,
	})
}

// GetMakesByCountry godoc
// @Summary      returns makes by country
// @Description  get makes by country
// @Tags         Make
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        country   query      string  true  "make search by country"
// @success 200 {object} dtos.SuccessResponseDto{data=[]models.Make} "desc"
// @Failure      400  {object}  dtos.FailedResponseDto	"request query validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"no make with the search country found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /makes/countries [get]
func GetMakesByCountry(context *gin.Context) {

	// Validate Request Query
	query := dtos.MakeCountryDto{}

	if err := context.BindQuery(&query); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	var makes []models.Make
	if foundMakes := config.DB.Where("country = ?", query.Country).Find(&makes).RowsAffected; foundMakes == 0 {
		exceptions.HandleNotFoundException(context, errors.New("no makes found for specified country"))
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "Makes with Country: " + query.Country,
		"data":       makes,
	})
}

// UpdateMake godoc
// @Summary      updates a make
// @Description  update make
// @Tags         Make
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Make ID(UUID)"
// @Param 		 data	body	dtos.CreateMakeDto	true	"Make Details JSON"
// @success 200 {object} dtos.SuccessResponseDto{data=models.Make} "make updated successfully"
// @Failure      400  {object}  dtos.FailedResponseDto	"request body/param validation errors or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"make with ID in request params not found"
// @Failure      409  {object}  dtos.FailedResponseDto	"another make with the same name in request body exists"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /makes/{id} [patch]
func UpdateMake(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}
	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Validate Request Body
	body := dtos.CreateMakeDto{}
	if err := context.BindJSON(&body); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Check if Make with specified name already exists
	var makeExists models.Make
	if err := config.DB.Where("name = ?", body.Name).Not("id = ?", params.ID).First(&makeExists).Error; err == nil {

		exceptions.HandleConflictException(context, "Make with name: "+body.Name+" already exists")
		return
	}

	var make models.Make
	result := config.DB.First(&make, "id = ?", params.ID)

	if result.Error != nil {

		exceptions.HandleNotFoundException(context, result.Error)
		return
	}

	config.DB.Model(&make).Updates(models.Make{Name: body.Name, Country: body.Country})

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Make Updated",
		"data":       make,
	})
}

// DeleteMake godoc
// @Summary      deletes a make
// @Description  delete make
// @Tags         Make
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Make ID(UUID)"
// @success 200 {object} dtos.SuccessResponseDto "make deleted successfully"
// @Failure      400  {object}  dtos.FailedResponseDto	"request params validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"make with ID in request params not found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /makes/{id} [delete]
func DeleteMake(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	result := config.DB.Delete(&models.Make{}, "id = ?", params.ID)

	if result.Error != nil {
		exceptions.HandleBadRequestException(context, result.Error)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Make Deleted",
	})
}
