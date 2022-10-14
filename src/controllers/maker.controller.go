package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/initializers"
	"github.com/greazleay/vehicle-api/src/models"
)

func CreateMaker(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateMaker{}
	var err = context.BindJSON(&body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"name":       "ValidationException",
			"status":     "failure",
			"statusCode": 400,
		})
		return
	}

	// Check if Maker with specified name already exists
	var makerExists models.Maker
	isError := initializers.DB.First(&makerExists, "name = ?", body.Name).Error

	if isError == nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error":      "Maker with name: " + body.Name + " already exists",
			"name":       "ConflictException",
			"status":     "failed",
			"statusCode": 409,
		})
		return
	}

	// Create new Maker
	newMaker := models.Maker{Name: body.Name}

	result := initializers.DB.Create(&newMaker)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":     "success",
		"statusCode": 201,
		"message":    "Maker Created",
		"data":       newMaker,
	})
}

func GetAllMakers(context *gin.Context) {

	var makers []models.Maker
	initializers.DB.Find(&makers)

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "All Makers",
		"data":       makers,
	})
}

func GetMakerByID(context *gin.Context) {

	id := context.Param("id")

	var maker models.Maker
	result := initializers.DB.First(&maker, "id = ?", id)

	if result.Error != nil {

		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":     "failure",
			"statusCode": 404,
			"name":       "NotFoundException",
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
	var err = context.BindJSON(&body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"name":       "ValidationException",
			"status":     "failure",
			"statusCode": 400,
		})
		return
	}

	// Check if Maker with specified name already exists
	var makerExists models.Maker
	isError := initializers.DB.First(&makerExists, "name = ?", body.Name).Error

	if isError == nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error":      "Maker with name: " + body.Name + " already exists",
			"name":       "ConflictException",
			"status":     "failed",
			"statusCode": 409,
		})
		return
	}

	id := context.Param("id")

	var maker models.Maker
	result := initializers.DB.First(&maker, "id = ?", id)

	if result.Error != nil {

		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":     "failure",
			"statusCode": 404,
			"name":       "NotFoundException",
			"error":      result.Error.Error(),
		})
		return
	}

	initializers.DB.Model(&maker).Updates(models.Maker{Name: body.Name})

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "Maker Updated",
		"data":       maker,
	})

}

func DeleteMaker(context *gin.Context) {

	id := context.Param("id")

	result := initializers.DB.Delete(&models.Maker{}, "id = ?", id)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": 200,
		"message":    "Maker Deleted",
	})
}
