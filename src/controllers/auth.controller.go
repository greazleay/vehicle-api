package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/greazleay/vehicle-api/src/config"
	"github.com/greazleay/vehicle-api/src/dtos"
	"github.com/greazleay/vehicle-api/src/exceptions"
	"github.com/greazleay/vehicle-api/src/helpers"
	"github.com/greazleay/vehicle-api/src/models"
	"github.com/greazleay/vehicle-api/src/services/auth"
)

// LoginUser godoc
// @Summary      login user with valid email and password combination
// @Description  login user
// @Tags         Auth
// @Security  BasicAuth
// @Accept       json
// @Produce      json
// @Param 		 data	body	dtos.LoginUserDto	true	"User Login Credentials JSON"
// @Success      200  {object}  dtos.SuccessResponseDto	"login successful"
// @Failure      400  {object}  dtos.FailedResponseDto	"request body validation errors"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid credentials"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /auth/login [post]
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

		exceptions.HandleUnauthorizedException(context, "Invalid Credentials")
		return
	}

	if invalidPasswordError := userExists.ValidatePassword(body.Password); invalidPasswordError != nil {

		exceptions.HandleUnauthorizedException(context, "Invalid Credentials")
		return

	}

	accessToken, err := auth.GenerateJwt(userExists.ID)
	if err != nil {

		exceptions.HandleInternalServerException(context)
		return
	}

	helpers.HandleOkResponse(context, "Login Successful", accessToken)
}
