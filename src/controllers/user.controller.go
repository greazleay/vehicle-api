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

// CreateUser godoc
// @Summary      registers a new user
// @Description  create user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 data	body	dtos.CreateUserDto	true	"New User Details JSON"
// @Success      201  {object}  dtos.SuccessResponseDto{data=models.User}	"user created successfully"
// @Failure      400  {object}  dtos.FailedResponseDto "request body validation error"
// @Failure      409  {object}  dtos.FailedResponseDto "another user with supplied email exists"
// @Failure      500  {object}  dtos.FailedResponseDto "unexpected internal server error"
// @Router       /users [post]
func CreateUser(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateUserDto{}

	if err := context.BindJSON(&body); err != nil {

		exceptions.HandleValidationException(context, err)
		return
	}

	// Check if Another User with specified email already exists
	var userExists models.User

	if err := config.DB.First(&userExists, "email = ?", body.Email).Error; err == nil {

		exceptions.HandleConflictException(context, "User with email: "+body.Email+" already exists")
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
		exceptions.HandleBadRequestException(context, result.Error)
		return
	}

	helpers.HandleCreatedResponse(context, "User Registered", newUser)
}

// GetAllUsers godoc
// @Summary      returns all users
// @Description  get all users
// @Tags         User
// @Security 	JWT
// @Accept       json
// @Produce      json
// @success 200 {object} dtos.SuccessResponseDto{data=[]models.User}	"all users returned"
// @Failure      400  {object}  dtos.FailedResponseDto	"token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /users [get]
func GetAllUsers(context *gin.Context) {

	var allUsers []models.User
	config.DB.Select("id", "email", "first_name", "last_name", "last_login", "created_at", "updated_at").Find(&allUsers)

	helpers.HandleOkResponse(context, "All Users", allUsers)
}

// GetUserByID godoc
// @Summary      returns a user by its 16 caharcter uuid
// @Description  get user by ID
// @Tags         User
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID(UUID)"
// @success 200 {object} dtos.SuccessResponseDto{data=models.User} "desc"
// @Failure      400  {object}  dtos.FailedResponseDto	"request param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"user with the specified ID not found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /users/{id} [get]
func GetUserByID(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", params.ID).Error; err != nil {

		exceptions.HandleNotFoundException(context, err)
		return
	}

	helpers.HandleOkResponse(context, "User with ID: "+params.ID, user)
}

// UpdateUser godoc
// @Summary      updates a user
// @Description  update user
// @Tags         User
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID(UUID)"
// @Param 		 data	body	dtos.UpdateUserDto	true	"User Details JSON"
// @success 200 {object} dtos.SuccessResponseDto{data=models.User}	"user updated successfully"
// @Failure      400  {object}  dtos.FailedResponseDto	"request body/param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"user with specified ID not found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /users/{id} [patch]
func UpdateUser(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}
	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Validate Request Body
	body := dtos.UpdateUserDto{}
	if err := context.BindJSON(&body); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Check if User with supplied ID exists
	var user models.User
	result := config.DB.First(&user, "id = ?", params.ID)

	if result.Error != nil {

		exceptions.HandleNotFoundException(context, result.Error)
		return
	}

	config.DB.Model(&user).Updates(models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
	})

	helpers.HandleOkResponse(context, "User Updated", user)

}

// DeleteUser godoc
// @Summary      deletes a user
// @Description  delete user
// @Tags         User
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID(UUID)"
// @success 200 {object} dtos.SuccessResponseDto	"user deleted suuceesfully"
// @Failure      400  {object}  dtos.FailedResponseDto	"request param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /users/{id} [delete]
func DeleteUser(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	result := config.DB.Delete(&models.User{}, "id = ?", params.ID)

	if result.Error != nil {
		exceptions.HandleBadRequestException(context, result.Error)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "User Deleted",
	})
}
