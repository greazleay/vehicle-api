package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/exceptions"
	"github.com/greazleay/vehicle-api/src/models"
	"gorm.io/gorm"
)

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

	context.JSON(http.StatusCreated, gin.H{
		"statusText": "success",
		"statusCode": 201,
		"message":    "Make Created",
		"data":       newMake,
	})
}

func GetAllMakes(context *gin.Context) {

	var makes []models.Make
	config.DB.Find(&makes)

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "All Makes",
		"data":       makes,
	})
}

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
