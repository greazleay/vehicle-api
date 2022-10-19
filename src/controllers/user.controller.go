package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/models"
)

func CreateUser(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateUserDto{}

	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

	// Check if Another User with specified email already exists
	var userExists models.User

	if err := config.DB.First(&userExists, "email = ?", body.Email).Error; err == nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"statusText": "failed",
			"statusCode": 409,
			"errorType":  "ConflictException",
			"error":      "User with email: " + body.Email + " already exists",
		})
		return
	}

	// Create new User
	newUser := models.User{
		Email:     body.Email,
		Password:  body.Password,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	result := config.DB.Create(&newUser)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"statusText": "success",
		"statusCode": 201,
		"message":    "User Registered",
		"data":       newUser,
	})
}

func GetAllUsers(context *gin.Context) {

	var users []models.User
	config.DB.Find(&users)

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "All Users",
		"data":       users,
	})
}

func GetUserByID(context *gin.Context) {

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

	var user models.User
	if err := config.DB.First(&user, "id = ?", params.ID).Error; err != nil {

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
		"message":    "User with ID: " + params.ID,
		"data":       user,
	})
}

func UpdateUser(context *gin.Context) {

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
	body := dtos.UpdateUserDto{}
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

	// Check if User with supplied ID exists
	var user models.User
	result := config.DB.First(&user, "id = ?", params.ID)

	if result.Error != nil {

		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"statusText": "failure",
			"statusCode": 404,
			"errorType":  "NotFoundException",
			"error":      result.Error.Error(),
		})
		return
	}

	config.DB.Model(&user).Updates(models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
	})

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Make Updated",
		"data":       user,
	})

}

func DeleteUser(context *gin.Context) {

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

	result := config.DB.Delete(&models.User{}, "id = ?", params.ID)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "User Deleted",
	})
}
