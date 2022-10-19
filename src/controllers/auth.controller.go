package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/middlewares/auth"
	"github.com/greazleay/vehicle-api/src/models"
)

func LoginUser(context *gin.Context) {

	// Validate Request Body
	body := dtos.LoginUserDto{}

	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusText": "failure",
			"statusCode": 400,
			"errorType":  "ValidationException",
			"error":      err.Error(),
		})
		return
	}

	// Check if User Exists
	var userExists models.User

	if err := config.DB.First(&userExists, "email = ?", body.Email).Error; err != nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"statusText": "failed",
			"statusCode": 401,
			"errorType":  "UnAuthorizedException",
			"error":      "Invalid Credentials",
		})
		return
	}

	if invalidPasswordError := userExists.ValidatePassword(body.Password); invalidPasswordError != nil {

		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"statusText": "failed",
			"statusCode": 401,
			"errorType":  "UnAuthorizedException",
			"error":      "Invalid Credentials",
		})
		return

	}

	tokenString, err := auth.GenerateJwt(userExists.ID)
	if err != nil {
		panic(err.Error())
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Login Successful",
		"data":       tokenString,
	})
}
