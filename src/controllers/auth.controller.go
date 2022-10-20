package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/exceptions"
	"github.com/greazleay/vehicle-api/src/models"
	"github.com/greazleay/vehicle-api/src/services/auth"
)

func LoginUser(context *gin.Context) {

	// Validate Request Body
	body := dtos.LoginUserDto{}

	if err := context.BindJSON(&body); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Check if User Exists
	var userExists models.User

	if err := config.DB.First(&userExists, "email = ?", body.Email).Error; err != nil {

		exceptions.HandleUnauthorizedException(context)
		return
	}

	if invalidPasswordError := userExists.ValidatePassword(body.Password); invalidPasswordError != nil {

		exceptions.HandleUnauthorizedException(context)
		return

	}

	accessToken, err := auth.GenerateJwt(userExists.ID)
	if err != nil {

		exceptions.HandleInternalServerException(context)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "Login Successful",
		"data":       accessToken,
	})
}
