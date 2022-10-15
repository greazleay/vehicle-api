package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/models"
)

func CreateMaker(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateMaker{}
	var err = context.BindJSON(&body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

	// Check if Maker with specified name already exists
	var makerExists models.Maker
	isError := config.DB.First(&makerExists, "name = ?", body.Name).Error

	if isError == nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"statusText": "failed",
			"statusCode": 409,
			"errorType":  "ConflictException",
			"error":      "Maker with name: " + body.Name + " already exists",
		})
		return
	}

	// Create new Maker
	newMaker := models.Maker{Name: body.Name}

	result := config.DB.Create(&newMaker)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"statusText": "success",
		"statusCode": 201,
		"message":    "Maker Created",
		"data":       newMaker,
	})
}

func GetAllMakers(context *gin.Context) {

	var makers []models.Maker
	config.DB.Find(&makers)

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "All Makers",
		"data":       makers,
	})
}

func GetMakerByID(context *gin.Context) {

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

	var maker models.Maker
	result := config.DB.First(&maker, "id = ?", params.ID)

	if result.Error != nil {

		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"statusText": "failure",
			"statusCode": 404,
			"errorType":  "NotFoundException",
			"error":      result.Error.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "Maker",
		"data":       maker,
	})
}

func UpdateMaker(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateMaker{}

	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

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

	// Check if Maker with specified name already exists
	var makerExists models.Maker

	if err := config.DB.First(&makerExists, "name = ?", body.Name).Error; err == nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"statusText": "failed",
			"statusCode": 409,
			"errorType":  "ConflictException",
			"error":      "Maker with name: " + body.Name + " already exists",
		})
		return
	}

	var maker models.Maker
	result := config.DB.First(&maker, "id = ?", params.ID)

	if result.Error != nil {

		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"statusText": "failure",
			"statusCode": 404,
			"errorType":  "NotFoundException",
			"error":      result.Error.Error(),
		})
		return
	}

	config.DB.Model(&maker).Updates(models.Maker{Name: body.Name})

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Maker Updated",
		"data":       maker,
	})

}

func DeleteMaker(context *gin.Context) {

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

	result := config.DB.Delete(&models.Maker{}, "id = ?", params.ID)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Maker Deleted",
	})
}
